package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if os.Getenv("SILENCE_OUTPUT") != "true" {
		fmt.Println("noop image used, exiting. (see github.com/rjocoleman/noop for more info)")
	}

	// Stay alive if NOOP_INFINITY=true
	if os.Getenv("NOOP_INFINITY") == "true" {
	sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
		<-sigChan
	}
}
