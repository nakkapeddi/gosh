package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func runCommand(commandString string) {
	defaultPath := []string{"/usr/local/bin", "/usr/bin", "/bin", "/usr/sbin", "/sbin", "/usr/local/go/bin"}
	commandFound := false
	slicedCommandWithArgs := stringSlicer(commandString)
	command := slicedCommandWithArgs[0]
	args := slicedCommandWithArgs[1:]

	if _, ok := builtins[command]; ok {
		runBuiltin(command, args...)
		return
	}

	for _, prefix := range defaultPath {
		fullCommandPath := prefix + "/" + slicedCommandWithArgs[0]
		if fileExists(fullCommandPath) {
			commandFound = true
			break
		}
	}

	if commandFound {
		stdout, stderr := execCommandString(command, args...)
		fmt.Printf(stdout)
		fmt.Printf(stderr)
	} else {
		Logger.Error().Msg("Couldn't find file")
	}

}

func execCommandString(cmd string, args ...string) (string, string) {
	cmder := exec.Command(cmd, args...)
	var out, sterr bytes.Buffer
	cmder.Stdout = &out
	cmder.Stderr = &sterr
	err := cmder.Run()
	if err != nil {
		Logger.Error().Err(err).Msg("")
	}
	return out.String(), sterr.String()
}
