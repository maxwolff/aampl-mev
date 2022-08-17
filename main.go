package main

import (
	"fmt"
	"os/exec"
)

func main() {
    cmd := exec.Command("forge", "t")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Print the output
    fmt.Println(string(stdout))
}