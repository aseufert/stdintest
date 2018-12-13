package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func execCommand(p string) {
	cmdArgs := strings.Fields(p)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...) //#nosec
	stdout, err := cmd.StdoutPipe()                             //#nosec
	cmd.Start()                                                 //#nosec
	//oneByte := make([]byte, 1000)
	for {
		//_, err := stdout.Read(oneByte)
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		r := bufio.NewReader(stdout)
		line, _, _ := r.ReadLine() //#nosec
		fmt.Printf("Line: %v\n", string(line))
	}

	cmd.Wait() //#nosec
}

func main() {
	execCommand("ls")
}
