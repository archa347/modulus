package cmd

import (
	"github.com/spf13/cobra"
)



func Execute() error {
	var rootCmd = &cobra.Command{
		Use: "modulus",
	}

	addDemoCommand(rootCmd)

	return rootCmd.Execute()
}
