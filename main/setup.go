package main

import (
	"fmt"
	"forward-router/core"
	"forward-router/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/ngaut/log"
)

func main() {
	fmt.Println("application has run.")
	logger.R().Println("hello my first exchange net application.")
	core.IStart()
	// handleSignal()
}
func handleSignal() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		sig := <-sigCh
		log.Infof("Got signal [%s] to exit.", sig)
		core.Stop()
	}()
}
