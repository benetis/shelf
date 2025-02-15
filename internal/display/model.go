package display

import (
	"github.com/benetis/shelf/internal"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keybindings []internal.Keybinding
	cursor      int
	debug       []string // what you see in the UI
	debugCh     <-chan string
	terminal    terminal
}

type terminal struct {
	width  int
	height int
}

func InitialModel(keybindings []internal.Keybinding, debugCh <-chan string) Model {
	return Model{
		keybindings: keybindings,
		cursor:      0,
		debug:       []string{},
		debugCh:     debugCh,
		terminal: terminal{
			width:  80,
			height: 24,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return listenForDebug(m.debugCh)
}

func listenForDebug(ch <-chan string) tea.Cmd {
	return func() tea.Msg {
		msg := <-ch
		return internal.DebugMsg(msg)
	}
}
