package main

import (
	"context"
	"fmt"
	"log"
    // "strconv"
	// crypto "crypto"
	"os/exec"

	"math/big"
    "strings"

	policy "ampl-arb/bindings/policy"
    arb "ampl-arb/bindings/arb"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	PolicyAddress = common.HexToAddress("0x1B228a749077b8e307C5856cE62Ef35d96Dca2ea")
    // 0x5fbdb2315678afecb367f032d93f642f64180aa3
	// PrivateKey, _ = crypto.HexToECDSA(os.Getenv("MEV_PRIVATE")) // Mev wallet private key
	RPC           = "http://localhost:8546" 
    PreemptRebaseSeconds = big.NewInt(30) // borrow from aave 30 seconds before rebase
    // os.Getenv("AVAX_WS_URL")
    // "https://mainnet.infura.io/v3/a92bf60569df4e02b242dbd8357fffac")
)

type AmplPolicy struct {
	ethClient *ethclient.Client
	policy *policy.Policy
    lastRebaseTimestampSec *big.Int
    minRebaseTimeIntervalSec *big.Int
    rebaseWindowOffsetSec *big.Int
    rebaseWindowLengthSec *big.Int
}

func (p *AmplPolicy) fetchState(){
    lastRebaseTimestampSec, err := p.policy.LastRebaseTimestampSec(nil)
    if err != nil {
        log.Fatal(err)
    }
    minRebaseTimeIntervalSec, err := p.policy.MinRebaseTimeIntervalSec(nil)
    if err != nil {
        log.Fatal(err)
    }

    rebaseWindowOffsetSec, err := p.policy.RebaseWindowOffsetSec(nil)
    if err != nil {
        log.Fatal(err)
    }

    rebaseWindowLengthSec, err := p.policy.RebaseWindowLengthSec(nil)

    p.lastRebaseTimestampSec = lastRebaseTimestampSec
    p.minRebaseTimeIntervalSec = minRebaseTimeIntervalSec
    p.rebaseWindowOffsetSec = rebaseWindowOffsetSec
    p.rebaseWindowLengthSec = rebaseWindowLengthSec
}

func (p *AmplPolicy) secondsUntilRebase(timestamp *big.Int) (*big.Int) {
    var lowerOffset = p.rebaseWindowOffsetSec
    var upperOffset *big.Int
    upperOffset.Add(lowerOffset, p.rebaseWindowLengthSec)

    var minRebaseTime *big.Int
    minRebaseTime.Add(p.lastRebaseTimestampSec, p.minRebaseTimeIntervalSec)

    if timestamp.Cmp(minRebaseTime) > 0 {
        minRebaseTime = timestamp
    }

    var remainder *big.Int
    remainder = remainder.Mod(minRebaseTime, p.minRebaseTimeIntervalSec)

    if remainder.Cmp(lowerOffset) < 0 {
        minRebaseTime := minRebaseTime.Add(minRebaseTime, lowerOffset)
        minRebaseTime = minRebaseTime.Sub(minRebaseTime, remainder)
    } else if remainder.Cmp(lowerOffset) > 0{
        minRebaseTime := minRebaseTime.Sub(minRebaseTime, remainder)
        minRebaseTime = minRebaseTime.Add(minRebaseTime, lowerOffset)
        minRebaseTime = minRebaseTime.Add(minRebaseTime, p.minRebaseTimeIntervalSec)
    }

    return minRebaseTime.Sub(minRebaseTime, timestamp)
}

func main() {
    fmt.Printf("started");

    out, err := exec.Command("./fork.sh").Output()
    
    client, err := ethclient.Dial(RPC)

    a, err := arb.NewArb("0x666d0c3da3dbc946d5128d06115bb4eed4595580", client)
    
    
    
    if err != nil {
        log.Fatal(err)
    }
    f := strings.Split(string(out),"\n")
    var bigI *big.Int;
    bigI.SetString(f[0], 10);
    // totalSupply := f[1]
    fmt.Printf(bigI.String());
    

}
func main2() {
    client, err := ethclient.Dial(RPC)

    if err != nil {
        log.Fatal(err)
    }

    p, err := policy.NewPolicy(PolicyAddress, client)

    if err != nil {
        log.Fatal(err)
    }

    policy := &AmplPolicy{
        ethClient: client, 
        policy: p,
    }

    
    for {
        policy.fetchState()

        header, err := client.HeaderByNumber(context.Background(), nil)
        if err != nil {
            log.Fatal(err)
        }
        secondsUntilRebase := policy.secondsUntilRebase(big.NewInt(int64(header.Time)))
        
        // if rebasing soon
        if secondsUntilRebase.Cmp(PreemptRebaseSeconds) < 0 {
            // check if profitable
            /*  do math in go
             * check math in foundry 
             * simulate in go. start anvil fork, read supply, call orchestrator, read supply, return
            */
            // policy.policy.Rebase()

        }
    }

}

