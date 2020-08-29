package main

type showDirectoryStack struct {
	builtin
}

var directoryStack []string

func (d *showDirectoryStack) setCommand(cmd string) {
	d.cmd = cmd
}

func (d *showDirectoryStack) setCommandArgs(args []string) {
	d.args = args
}

func (d *showDirectoryStack) runCommand() {
	Logger.Info().Msgf("%s", directoryStack)
}
