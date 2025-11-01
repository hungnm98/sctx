package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(lsCmd) }

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, _ := os.UserHomeDir()
		dir := filepath.Join(home, ".sctx", "profiles")
		entries, err := os.ReadDir(dir)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("(no profiles)")
				return nil
			}
			return err
		}
		current := ""
		if data, err := os.ReadFile(filepath.Join(home, ".sctx", "current")); err == nil {
			current = string(data)
		}

		for _, e := range entries {
			name := e.Name()
			if name == current {
				fmt.Printf("* %s\n", name)
			} else {
				fmt.Printf("  %s\n", name)
			}
		}
		return nil
	},
}
