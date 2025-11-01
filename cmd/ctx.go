package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(ctxCmd) }

var ctxCmd = &cobra.Command{
	Use:   "ctx",
	Short: "Interactive profile selector",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 🚫 Nếu đang trong subshell, chặn đổi profile
		if os.Getenv("SCTX_ACTIVE") == "1" {
			current := os.Getenv("SCTX_PROFILE")
			fmt.Printf("⚠️  You are currently inside profile %q.\n", current)
			fmt.Println("👉 Please type 'exit' to leave the subshell before switching to another profile.")
			return nil
		}

		home, _ := os.UserHomeDir()
		dir := filepath.Join(home, ".sctx", "profiles")
		entries, err := os.ReadDir(dir)
		if err != nil {
			return err
		}
		if len(entries) == 0 {
			fmt.Println("no profiles found")
			return nil
		}

		var items []string
		for _, e := range entries {
			items = append(items, e.Name())
		}

		prompt := promptui.Select{
			Label: "Select profile",
			Items: items,
		}

		_, result, err := prompt.Run()
		if err != nil {
			return err
		}

		return useCmd.RunE(cmd, []string{result})
	},
}
