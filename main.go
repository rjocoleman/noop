package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Getenv("SILENCE_OUTPUT") != "true" {
		fmt.Println("noop image used, exiting. (see github.com/rjocoleman/noop for more info)")
	}
}
