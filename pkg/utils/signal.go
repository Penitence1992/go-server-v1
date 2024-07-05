package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func CreateInterruptChan() <-chan os.Signal {
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	return sign
}
