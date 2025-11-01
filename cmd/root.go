package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sctx",
	Short: "Simple shell context switcher",
	Long:  "sctx lets you create, edit, list and switch between shell profiles stored in ~/.sctx/profiles",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
