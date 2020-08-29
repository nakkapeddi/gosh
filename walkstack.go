package main

type walkDirectoryStack struct {
	builtin
}

func (wds *walkDirectoryStack) setCommand(cmd string) {
	wds.cmd = cmd
}

func (wds *walkDirectoryStack) setCommandArgs(args []string) {
	wds.args = args
}

func (wds *walkDirectoryStack) runCommand() {
	if len(directoryStack) == 0 {
		return
	}
	switch wds.cmd {
	case "b":
		if pos == 0 {
			Logger.Debug().Msg("At bottom of directory stack")
		} else {
			pos--
			cd := changeDirectory{}
			cd.changeDirectory(directoryStack[pos])
		}
	case "a":
		if pos == len(directoryStack)-1 {
			Logger.Debug().Msg("At top of directory stack")
		} else {
			pos++
			cd := changeDirectory{}
			cd.changeDirectory(directoryStack[pos])
		}

	}

}
