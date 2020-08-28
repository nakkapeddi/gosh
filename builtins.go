package main

import "os"

var builtins = map[string]interface{}{
	"cd": changeDirectory,
	"d":  showDirectoryStack,
}

func runBuiltin(commandString string, commandArgs ...string) {
	builtins[commandString].(func(...string))(commandArgs...)
}

var directoryStack []string

func changeDirectory(dirs ...string) {
	Logger.Debug().Msgf("cd: %s\n", dirs)
	if len(dirs) == 0 {
		return
	}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); err != nil {
			Logger.Error().Err(err).Msg("")
			return
		}
		directoryStack = append(directoryStack, dir)
	}

	topElem := len(directoryStack) - 1
	err := os.Chdir(directoryStack[topElem])
	if err != nil {
		Logger.Error().Err(err).Msg("")
	}
}

func showDirectoryStack(...string) {
	Logger.Info().Msgf("%s", directoryStack)
}
