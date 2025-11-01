package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(unsetCmd) }

var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset current profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, _ := os.UserHomeDir()
		current := filepath.Join(home, ".sctx", "current")
		if err := os.Remove(current); err != nil && !os.IsNotExist(err) {
			return err
		}
		fmt.Println("✅ unset current profile")
		return nil
	},
}
