package main

import (
	"github.com/archa347/modulus/command/scheduler"
	"os"
)

func main() {
	os.Exit(runCommand())
}

func runCommand() int {
	scheduler.Run(os.Args[1:])

	return 0
}
