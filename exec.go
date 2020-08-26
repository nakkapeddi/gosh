package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func runCommand(commandString string) {
	defaultPath := []string{"/usr/local/bin", "/usr/bin", "/bin", "/usr/sbin", "/sbin", "/usr/local/go/bin"}
	commandFound := false
	slicedCommandWithArgs := stringSlicer(commandString)

	for _, prefix := range defaultPath {
		fullCommandPath := prefix + "/" + slicedCommandWithArgs[0]
		if fileExists(fullCommandPath) {
			commandFound = true
			break
		}
	}

	if commandFound {
		execCommandString(slicedCommandWithArgs[0], slicedCommandWithArgs)
	} else {
		Logger.Error().Msg("Couldn't find file")
	}

}

func execCommandString(cmd string, args []string) {
	cmder := exec.Command(cmd, args[1:]...)
	var out bytes.Buffer
	cmder.Stdout = &out
	err := cmder.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String())
}
