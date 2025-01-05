package utils

import (
	"os"
	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger(level string) {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	Log = zerolog.New(os.Stdout).With().
		Timestamp().
		Logger().Level(logLevel)
}