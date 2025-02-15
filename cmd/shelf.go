package cmd

import (
	"fmt"
	"github.com/benetis/shelf/internal"
	"github.com/benetis/shelf/internal/display"
	"github.com/benetis/shelf/internal/hammerparser"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shelf",
	Short: "Displays configured shortcuts for various applications.",
	Run: func(cmd *cobra.Command, args []string) {
		var keybindings []internal.Keybinding

		// hammerspoon
		defaultPath := "~/.hammerspoon"
		keybindings = slices.Concat(keybindings, hammerparser.Parse(defaultPath, true))

		p := tea.NewProgram(display.InitialModel(keybindings))

		_, err := p.Run()
		if err != nil {
			panic(fmt.Errorf("cannot run tea program: %w", err))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
