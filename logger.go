package main

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func setupLogger() {
	rootLogger := zerolog.New(os.Stderr).With().Caller().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Logger = rootLogger
}
