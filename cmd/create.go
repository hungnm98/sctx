package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(createCmd) }

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		home, _ := os.UserHomeDir()
		dir := filepath.Join(home, ".sctx", "profiles")
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}

		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("profile %q already exists", name)
		}

		content := fmt.Sprintf(`# ~/.sctx/profiles/%s
# put your shell config here

# Profile environment variable
export SCTX_PROFILE=%s

# Colors
CYAN='\033[1;36m'
NC='\033[0m'

# Set PS1 to show profile prefix (only for interactive shells)
if [[ $- == *i* ]]; then
  export PS1="[\033[1;36m%s\033[0m] $PS1"
fi

# Highlight when switching context
echo -e "\033[1;36m[%s]\033[0m switched to %s environment"
`, name, name, name, name, name)

		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return err
		}

		fmt.Printf("✅ created profile %s at %s\n", name, path)
		return nil
	},
}
