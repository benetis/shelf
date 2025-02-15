package display

import (
	"github.com/benetis/shelf/internal"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keybindings  []internal.Keybinding
	cursor       int
	debugEnabled bool
	debug        []string // what you see in the UI
	debugCh      <-chan string
	terminal     terminal
}

type terminal struct {
	width  int
	height int
}

func InitialModel(keybindings []internal.Keybinding, debugEnabled bool, debugCh <-chan string) Model {
	return Model{
		keybindings:  keybindings,
		cursor:       0,
		debugEnabled: debugEnabled,
		debug:        []string{},
		debugCh:      debugCh,
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
