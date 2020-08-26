package main

func main() {
	setupLogger()
	for {
		SignalHandler()
		cmd, err := prompt()
		if err != nil {
			Logger.Debug().Err(err).Msg("")
		} else {
			runCommand(cmd)
		}

	}
}
