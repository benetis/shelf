package display

import (
	"github.com/benetis/shelf/internal"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keybindings []internal.Keybinding
	cursor      int

	terminal terminal
}

type terminal struct {
	width  int
	height int
}

func InitialModel(keybindings []internal.Keybinding) Model {
	return Model{
		keybindings: keybindings,
		cursor:      0,
		terminal: terminal{
			width:  80,
			height: 24,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
