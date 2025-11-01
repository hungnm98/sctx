package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sctx",
	Short: "Simple shell context switcher",
	Long:  "sctx lets you create, edit, list and switch between shell profiles stored in ~/.sctx/profiles",
}

func Execute() {
	args := os.Args[1:]
	if len(args) == 0 || (len(args) == 1 && !strings.HasPrefix(args[0], "-")) {
		os.Args = append(os.Args, "ctx")
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
