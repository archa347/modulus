package main

import (
	"github.com/archa347/modulus/cmd"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Errorf("Error %s", err.Error())
		os.Exit(1)
	}
}
