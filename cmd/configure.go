/*
 * Copyright © 2024 Carlos Mendez <carlos@hadaelectronics.com>
 */
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cjairm/devenv/tools"
	"github.com/spf13/cobra"
)

var (
	All          bool
	OhMyZsh      bool
	Nvim         bool
	Tmux         bool
	WindowManage bool
	configureCmd = &cobra.Command{
		Use:   "configure",
		Short: "Sets new configuration",
		Long: `
You can configure them them separately by

nvim
yabai
skhd
	`,
		Run: configureRun,
	}
)

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().BoolVarP(&All, "all", "a", false, "Configures all available apps")
	configureCmd.Flags().BoolVarP(&OhMyZsh, "oh-my-zsh", "o", false, "Add custom aliases")
	configureCmd.Flags().BoolVarP(&Nvim, "nvim", "n", false, "Configures kickstart.nvim")
	configureCmd.Flags().BoolVarP(&Tmux, "tmux", "t", false, "Configures tmux basics")
	configureCmd.Flags().
		BoolVarP(&WindowManage, "window-manage", "w", false, "Configures yabai and skhd")
}

func configureRun(cmd *cobra.Command, args []string) {
	projectPath := tools.GetProjectPath()

	fmt.Printf("\n")
	configureOhMyZsh(projectPath)
	configureNvim(projectPath)
	configureTmux(projectPath)
	configureWindowManagement(projectPath)
	fmt.Printf("\n")
}

func configureOhMyZsh(projectPath string) {
	if OhMyZsh || All {
		currentZshFile, err := tools.ReadAllContent(".zshrc")
		if err != nil {
			return
		}
		err = tools.ReadFile(
			fmt.Sprintf("%soh_my_zsh/.extension_zshrc", projectPath),
			func(line string) {
				if line != "" && !strings.Contains(currentZshFile, line) &&
					!strings.HasPrefix(line, "#") {
					_ = tools.AppendStringToFile(".zshrc", line)
				}
			},
		)
		if err != nil {
			return
		}
		tools.HighlightLine("zsh configured")
	}
}

func configureNvim(projectPath string) {
	if Nvim || All {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home dir:", err)
			return
		}
		dst := filepath.Dir(fmt.Sprintf("%s/%s", homeDir, ".config/nvim/"))
		err = tools.CopyDirectory(fmt.Sprintf("%snvim", projectPath), dst)
		if err != nil {
			return
		}
		tools.HighlightLine("nvim configured")
	}
}

func configureTmux(projectPath string) {
	if Tmux || All {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home dir:", err)
			return
		}
		err = tools.CopyFile(
			fmt.Sprintf("%stmux/.tmux.conf", projectPath),
			fmt.Sprintf("%s/%s", homeDir, ".tmux.conf"),
		)
		if err != nil {
			return
		}
		tools.HighlightLine("tmux configured")
	}
}

func configureWindowManagement(projectPath string) {
	if WindowManage || All {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home dir:", err)
			return
		}
		err = tools.CopyDirectory(
			fmt.Sprintf("%sskhd", projectPath),
			fmt.Sprintf("%s/%s", homeDir, ".config/skhd"),
		)
		if err != nil {
			return
		}
		tools.HighlightLine("skhd configured")
		err = tools.CopyDirectory(
			fmt.Sprintf("%syabai", projectPath),
			fmt.Sprintf("%s/%s", homeDir, ".config/yabai"),
		)
		if err != nil {
			return
		}
		tools.HighlightLine("yabai configured")
	}
}
