package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func prompt() (string, error) {
	var userInputCommand string
	fmt.Printf("\342\232\241 ")
	stdin := bufio.NewScanner(os.Stdin)
	scanner := stdin.Scan()
	if scanner {
		rawText := stdin.Text()
		sanitized := strings.TrimLeft(rawText, "\t \n")
		userInputCommand = sanitized
	}
	eof := !scanner && stdin.Err() == nil
	if eof {
		Logger.Debug().Msg("Caught CTRL+D, exiting...")
		os.Exit(0)
	}
	if len(userInputCommand) == 0 {
		return "", errors.New("EMPTYMESSAGE")
	}
	return userInputCommand, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func stringSlicer(whiteSpaceDelimittedString string) []string {
	substringSlice := strings.Fields(whiteSpaceDelimittedString)
	return substringSlice
}
