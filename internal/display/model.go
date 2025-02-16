package display

import (
	"github.com/benetis/shelf/internal"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const defaultListHeight = 18

type Model struct {
	list         list.Model
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
	var items []list.Item
	for _, kb := range keybindings {
		items = append(items, keybindingItem{kb: kb})
	}

	l := list.New(items, list.NewDefaultDelegate(), 80, defaultListHeight)
	l.Title = "Shelf - Custom keybindings in one place"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return Model{
		list:         l,
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
