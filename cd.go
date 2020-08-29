package main

import "os"

type changeDirectory struct {
	builtin
}

func (cd *changeDirectory) setCommand(cmd string) {
	cd.cmd = cmd
}

func (cd *changeDirectory) setCommandArgs(args []string) {
	cd.args = args
}

func (cd *changeDirectory) runCommand() {
	Logger.Debug().Msgf("cd: %s\n", cd.args)
	if len(cd.args) == 0 {
		return
	}
	for _, dir := range cd.args {
		if _, err := os.Stat(dir); err != nil {
			Logger.Error().Err(err).Msg("")
			return
		}
		directoryStack = append(directoryStack, dir)
	}

	topElem := len(directoryStack) - 1
	err := cd.changeDirectory(directoryStack[topElem])
	if err != nil {
		Logger.Error().Err(err).Msg("")
	}
	pos = len(directoryStack) - 1
}

func (cd *changeDirectory) changeDirectory(dir string) error {
	err := os.Chdir(dir)
	return err
}
