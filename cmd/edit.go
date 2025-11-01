package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Editor string `yaml:"editor"`
}

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit [profile]",
	Short: "Edit a profile in your preferred editor",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		home, _ := os.UserHomeDir()
		sctxDir := filepath.Join(home, ".sctx")
		profileDir := filepath.Join(sctxDir, "profiles")
		if err := os.MkdirAll(profileDir, 0o755); err != nil {
			return err
		}
		profilePath := filepath.Join(profileDir, name)

		if _, err := os.Stat(profilePath); os.IsNotExist(err) {
			// ask create
			fmt.Printf("profile %q not found. create it? [y/N]: ", name)
			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			if strings.ToLower(line) != "y" {
				return nil
			}
			if err := os.WriteFile(profilePath, []byte(fmt.Sprintf("# ~/.sctx/profiles/%s\n", name)), 0o644); err != nil {
				return err
			}
		}

		cfgPath := filepath.Join(sctxDir, "config.yaml")
		cfg := loadConfig(cfgPath)
		editor := resolveEditor(cfg)
		if cfg.Editor != editor {
			cfg.Editor = editor
			saveConfig(cfgPath, cfg)
		}

		c := exec.Command(editor, profilePath)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	},
}

func loadConfig(path string) Config {
	var cfg Config
	data, err := os.ReadFile(path)
	if err == nil {
		_ = yaml.Unmarshal(data, &cfg)
	}
	return cfg
}

func saveConfig(path string, cfg Config) {
	data, _ := yaml.Marshal(cfg)
	os.WriteFile(path, data, 0o644)
}

func resolveEditor(cfg Config) string {
	if cfg.Editor != "" {
		return cfg.Editor
	}
	if env := os.Getenv("EDITOR"); env != "" {
		return env
	}
	candidates := []string{"vim", "nano", "code", "micro", "emacs", "vi"}
	available := []string{}
	for _, c := range candidates {
		if _, err := exec.LookPath(c); err == nil {
			available = append(available, c)
		}
	}
	if len(available) == 0 {
		return "vi"
	}
	if len(available) == 1 {
		return available[0]
	}
	fmt.Println("No editor configured. Select one:")
	for i, e := range available {
		fmt.Printf("  %d) %s\n", i+1, e)
	}
	fmt.Print("Choice [1-", len(available), "]: ")
	var choice int
	fmt.Scanln(&choice)
	if choice < 1 || choice > len(available) {
		fmt.Println("invalid choice, using", available[0])
		return available[0]
	}
	return available[choice-1]
}
