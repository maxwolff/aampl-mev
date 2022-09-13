package main

import (
	"context"
	"fmt"
	"log"

	lendingPoolBinding "ampl-arb/bindings/ilendingpool"
	viewBinding "ampl-arb/bindings/view"
	ecdsa "crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var (
	MinProfit            = big.NewInt(5e17)                                                  // .5eth
	PreemptRebaseSeconds = big.NewInt(100)                                                   // borrow from aave 100 seconds before rebase
	viewAddr             = common.HexToAddress("0x3fdc08d815cc4ed3b7f69ee246716f2c8bcd6b07") // default deploy addr of first default anvil addr
	amplAddr             = common.HexToAddress("0xD46bA6D942050d489DBd938a2C909A5d5039A161")
	lendingPool          = common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")
	policy               = common.HexToAddress("0x1B228a749077b8e307C5856cE62Ef35d96Dca2ea")
)

type Bot struct {
	Client     *ethclient.Client
	WsClient   *ethclient.Client
	SigningKey *ecdsa.PrivateKey // flashbots key
	PrivateKey *ecdsa.PrivateKey
	EOA        *common.Address // EOA holding funds
	Addresses  *ContractAddresses
}

type ContractAddresses struct {
	ViewContractAddr *common.Address // address of view contract
	AmplAddr         *common.Address // token address
	LendingPool      *common.Address // aave entry point
	Policy           *common.Address // aampl rebase contract
}

func main() {
	addrs := &ContractAddresses{
		ViewContractAddr: &viewAddr,
		AmplAddr:         &amplAddr,
		LendingPool:      &lendingPool,
		Policy:           &policy,
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	privateKey := os.Getenv("BOT_PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("BOT_PRIVATE_KEY not set")
	}

	signingKey := os.Getenv("FLASHBOTS_SIGNING_KEY")
	if signingKey == "" {
		log.Fatal("FLASHBOTS_SIGNING_KEY not set")
	}
	providerURL := os.Getenv("PROVIDER_URL")
	if providerURL == "" {
		log.Fatal("PROVIDER_URL not set")
	}

	wsproviderURL := os.Getenv("WS_PROVIDER_URL")
	if providerURL == "" {
		log.Fatal("PROVIDER_URL not set")
	}

	bot, err := NewBot(signingKey, privateKey, providerURL, wsproviderURL, addrs)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %s\n", err)
	}

	view, err := viewBinding.NewView(*bot.Addresses.ViewContractAddr, bot.Client)

	if err != nil {
		log.Fatalf("Failed to initialize view view client: %s\n", err)
	}

	eventSignature := []byte("LogRebase(uint256,uint256,uint256,int256,uint256)")
	rebaseSig := crypto.Keccak256Hash(eventSignature)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{*bot.Addresses.Policy},
		Topics:    [][]common.Hash{{rebaseSig}},
	}

	logs := make(chan types.Log)
	sub, err := bot.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// if rebase, repay
	go func() {
		select {
		case err := <-sub.Err():
			log.Fatal("sub err", err)
		case rebaseLog := <-logs:
			fmt.Println("saw rebase", rebaseLog)
			debt, err := view.UserDebt(&bind.CallOpts{}, *bot.EOA)
			if err != nil {
				fmt.Println("Failed to get user debt", err)
			} else {
				if debt != big.NewInt(0) {
					//repay full
					bot.repayAave()
				}
			}
		}
	}()

	// if before rebase, borrow
	for {
		secondsUntilRebase, err := bot.getSecondsUntilRebase()
		fmt.Println(secondsUntilRebase, err)
		if err != nil {
			fmt.Println(err)
		} else {
			// if rebasing soon
			expectedProfit, borrowAmount, err := bot.getExpectedProfit()
			if err != nil {
				fmt.Println(err)
			}
			debt, err := view.UserDebt(&bind.CallOpts{}, *bot.EOA)
			if err != nil {
				fmt.Println("Failed to get user debt", err)
			}

			// if right before rebase, expected to be profitable, and not already borrowed
			if secondsUntilRebase.Cmp(PreemptRebaseSeconds) < 1 && expectedProfit.Cmp(MinProfit) == 1 && debt == big.NewInt(0) {
				bot.borrowAave(borrowAmount)
			}
		}

		// sleep 10 sec before checking again
		fmt.Println("... ")
		time.Sleep(10 * 1e9)
	}
}

func (bot *Bot) getSecondsUntilRebase() (*big.Int, error) {
	abi, err := abi.JSON(strings.NewReader(viewBinding.ViewABI))
	if err != nil {
		return nil, err
	}
	data, err := abi.Pack("secondsUntilRebase")

	if err != nil {
		return nil, err
	}

	callData := &ethereum.CallMsg{
		From: *bot.EOA,
		To:   bot.Addresses.ViewContractAddr,
		Data: data,
	}

	rawCallResult, err := bot.Client.CallContract(context.Background(), *callData, nil)
	if err != nil {
		return nil, err
	}

	returnValues, err := abi.Unpack("secondsUntilRebase", rawCallResult)
	if err != nil {
		return nil, err
	}

	return returnValues[0].(*big.Int), nil
}

func (bot *Bot) getExpectedProfit() (*big.Int, *big.Int, error) {
	abi, err := abi.JSON(strings.NewReader(viewBinding.ViewABI))
	if err != nil {
		return nil, nil, err
	}
	// data, err := abi.Pack("expectedProfit", common.HexToAddress("0xb394e0a990a577f277ddb65547b6399898d78051"));
	data, err := abi.Pack("expectedProfit", *bot.EOA)
	if err != nil {
		return nil, nil, err
	}
	callData := &ethereum.CallMsg{
		From: *bot.EOA,
		To:   bot.Addresses.ViewContractAddr,
		Data: data,
	}

	rawCallResult, err := bot.Client.CallContract(context.Background(), *callData, nil)
	if err != nil {
		log.Fatalf("Failed to call contract: %s\n", err)
	}

	returnValues, err := abi.Unpack("expectedProfit", rawCallResult)
	expectedProfit := returnValues[0]
	borrowAmount := returnValues[1]
	if err != nil {
		return nil, nil, err
	}

	return expectedProfit.(*big.Int), borrowAmount.(*big.Int), nil
}

func (bot *Bot) repayAave() error {
	abi, err := abi.JSON(strings.NewReader(lendingPoolBinding.IlendingpoolABI))
	if err != nil {
		return err
	}
	// repay max uint to repay all. (1 << 256) - 1
	data, err := abi.Pack("repay", bot.Addresses.AmplAddr, maxUint256(), 2, *bot.EOA)
	if err != nil {
		return err
	}
	callData := &ethereum.CallMsg{
		From: *bot.EOA,
		To:   bot.Addresses.LendingPool,
		Data: data,
	}
	signedTx, err := bot.buildTransaction(bot.Addresses.LendingPool, big.NewInt(0), *callData)
	if err != nil {
		return err
	}
	return bot.sendPrivateFlashbotsTransaction(signedTx)
}

func (bot *Bot) borrowAave(amount *big.Int) error {
	abi, err := abi.JSON(strings.NewReader(lendingPoolBinding.IlendingpoolABI))
	if err != nil {
		return err
	}
	data, err := abi.Pack("borrow", bot.Addresses.AmplAddr, amount, 2, 0, *bot.EOA)
	if err != nil {
		return err
	}
	callData := &ethereum.CallMsg{
		From: *bot.EOA,
		To:   bot.Addresses.LendingPool,
		Data: data,
	}
	signedTx, err := bot.buildTransaction(bot.Addresses.LendingPool, big.NewInt(0), *callData)
	if err != nil {
		return err
	}
	return bot.sendPrivateFlashbotsTransaction(signedTx)
}

func (bot *Bot) buildTransaction(to *common.Address, value *big.Int, callData ethereum.CallMsg) (*types.Transaction, error) {

	nonce, err := bot.Client.PendingNonceAt(context.Background(), *bot.EOA)
	if err != nil {
		return nil, err
	}

	gasPrice, err := bot.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasLimit, err := bot.Client.EstimateGas(context.Background(), callData)
	if err != nil {
		log.Printf("Caught error estimating gas: %s\n", err)
		gasLimit = 400000
	}

	chainID, err := bot.Client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit * 2,
		To:       to,
		Value:    value,
		Data:     callData.Data,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), bot.PrivateKey)
	if err != nil {
		fmt.Println("Error signing transaction: ", err)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (bot *Bot) sendPrivateFlashbotsTransaction(signedTx *types.Transaction) error {
	data, err := signedTx.MarshalBinary()
	if err != nil {
		return err
	}
	txStr := hexutil.Encode(data)

	opts := flashbotsrpc.FlashbotsSendPrivateTransactionRequest{
		Tx: txStr,
		Preferences: &flashbotsrpc.FlashbotsPrivateTxPreferences{
			Fast: true,
		},
	}

	rpc := flashbotsrpc.New("https://relay.flashbots.net")

	result, err := rpc.FlashbotsSendPrivateTransaction(bot.PrivateKey, opts)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) { // user/tx error, rather than JSON or network error
			fmt.Println(err.Error())
		} else {
			fmt.Printf("error: %+v\n", err)
		}
	}

	// Print result
	fmt.Printf("%+v\n", result)

	return nil
}

func NewBot(signingKey string, privateKey string, providerURL string, wsProviderUrl string, addrs *ContractAddresses) (*Bot, error) {
	var err error
	bot := Bot{}

	bot.Client, err = ethclient.Dial(providerURL)
	if err != nil {
		return nil, err
	}

	bot.WsClient, err = ethclient.Dial(wsProviderUrl)
	if err != nil {
		return nil, err
	}

	bot.SigningKey, err = crypto.HexToECDSA(signingKey)
	if err != nil {
		return nil, err
	}

	bot.PrivateKey, bot.EOA, err = wallet(privateKey)
	if err != nil {
		return nil, err
	}

	bot.Addresses = addrs

	return &bot, nil
}

func wallet(private string) (*ecdsa.PrivateKey, *common.Address, error) {

	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return nil, nil, err
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	return privateKey, &fromAddress, nil
}

func maxUint256() *big.Int {
	var maxUint *big.Int
	maxUint.Lsh(big.NewInt(1), 256)
	maxUint.Sub(maxUint, big.NewInt(1))
	return maxUint
}
