package main

import (
	"fmt"
	"os"
	"os/signal"
)

func SignalHandler() {
	c := make(chan os.Signal)
	signal.Notify(c)
	go func() {
		s := <-c
		if s == os.Interrupt {
			fmt.Println("\rCtrl+C caught, exiting!")
			os.Exit(1)
		} else {
			Logger.Debug().Msgf("Signal caught: %s", s)
		}

	}()
}
