/*
 * Copyright © 2024 Carlos Mendez <carlos@hadaelectronics.com>
 */
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devenv",
	Short: "Set up your dev enviroment",
	Long: `
Commands that helps setting up dev tools

nvim: Kickstarter

yabai: ------

skdh: ----
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
