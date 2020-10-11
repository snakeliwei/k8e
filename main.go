package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xiaods/k8e/pkg/cli/cmds"
	"github.com/xiaods/k8e/pkg/cli/master"
	"github.com/xiaods/k8e/pkg/version"
)

func main() {
	cmdVersion := version.MakeVersion()

	printk8eASCIIArt := version.PrintK8eASCIIArt

	var rootCmd = &cobra.Command{
		Use: "k8e",
		Run: func(cmd *cobra.Command, args []string) {
			printk8eASCIIArt()
			cmd.Help()
		},
	}

	rootCmd.AddCommand(cmdVersion)
	rootCmd.AddCommand(cmds.NewMasterCommand(master.Run))

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
