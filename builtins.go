package main

type RunCommandable interface {
	setCommand(cmd string)
	setCommandArgs(args []string)
	runCommand()
}

type builtin struct {
	cmd  string
	args []string
}

var builtins = map[string]RunCommandable{
	"cd": &changeDirectory{},
	"d":  &showDirectoryStack{},
}

func runBuiltin(commandString string, commandArgs ...string) {
	builtins[commandString].setCommand(commandString)
	builtins[commandString].setCommandArgs(commandArgs)
	builtins[commandString].runCommand()
}
