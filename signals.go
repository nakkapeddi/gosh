package main

import (
	"os"
	"os/signal"
	"syscall"
)

func SignalHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-c
		Logger.Debug().Msgf("\rSignal caught: %s", s)
		switch s {
		case os.Interrupt:
			Logger.Warn().Msg("\rCaught interrupt")
		case syscall.SIGTERM:
			Logger.Warn().Msg("\rHandling SIGTERM, exiting gracefully...")
			os.Exit(0)
		}
	}()
}
