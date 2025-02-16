package cmd

import (
	"fmt"
	"github.com/benetis/shelf/internal"
	"github.com/benetis/shelf/internal/display"
	"github.com/benetis/shelf/internal/flashspace"
	"github.com/benetis/shelf/internal/hammerspoon"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

var debugFlag bool

var rootCmd = &cobra.Command{
	Use:   "shelf",
	Short: "Displays configured shortcuts for various applications.",
	Run: func(cmd *cobra.Command, args []string) {
		var keybindings []internal.Keybinding

		debugWriter := internal.NewDebugWriter()
		log.SetOutput(debugWriter)

		// hammerspoon
		defaultPath := "~/.hammerspoon"
		keybindings = slices.Concat(keybindings, hammerspoon.Parse(defaultPath, debugFlag))

		// flashspace
		defaultPath = "~/.config/flashspace/profiles.json"
		keybindings = slices.Concat(keybindings, flashspace.Parse(defaultPath, debugFlag))

		p := tea.NewProgram(display.InitialModel(keybindings, debugFlag, debugWriter.Channel()))

		_, err := p.Run()
		if err != nil {
			panic(fmt.Errorf("cannot run tea program: %w", err))
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Enable debug output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
