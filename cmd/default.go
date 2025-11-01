package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(defaultCmd) }

var defaultCmd = &cobra.Command{
	Use:   "default [name]",
	Short: "Set default profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		home, _ := os.UserHomeDir()
		profilePath := filepath.Join(home, ".sctx", "profiles", name)
		if _, err := os.Stat(profilePath); os.IsNotExist(err) {
			return fmt.Errorf("profile %q not found", name)
		}
		defaultFile := filepath.Join(home, ".sctx", "default")
		if err := os.WriteFile(defaultFile, []byte(name), 0o644); err != nil {
			return err
		}
		fmt.Printf("✅ default profile set to: %s\n", name)
		return nil
	},
}
