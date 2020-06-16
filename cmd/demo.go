package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func addDemoCommand(rootCommand *cobra.Command) {
	demoCommand := &cobra.Command{
		Use: "demo",
		Run: func(cmd *cobra.Command, args []string) {
			runDemo()
		},
	}

	rootCommand.AddCommand(demoCommand)
}

func runDemo() {
	log.Info("Running demo")
}