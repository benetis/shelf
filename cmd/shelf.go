package cmd

import (
	"github.com/benetis/shelf/internal/hammerparser"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shelf",
	Short: "Displays configured shortcuts for various applications.",
	Run: func(cmd *cobra.Command, args []string) {
		defaultPath := "~/.hammerspoon"
		hammerparser.Parse(defaultPath, true)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
