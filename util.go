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
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		rawText := scanner.Text()
		sanitized := strings.TrimLeft(rawText, "\t \n")
		userInputCommand = sanitized
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
