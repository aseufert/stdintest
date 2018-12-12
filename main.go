package main

import (
	"fmt"
	"os"
	"os/exec"
)

func execCommand(p string, a ...string) {
	fmt.Println("Running...")
	cmd := exec.Command(p, a...) // #nosec
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	execCommand("idevicesyslog")
}
