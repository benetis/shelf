package display

import (
	"github.com/benetis/shelf/internal"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminal.width = msg.Width
		m.terminal.height = msg.Height

		listHeight := defaultListHeight
		if m.debugEnabled { // If debug is enabled, make room for the debug console.
			listHeight = m.terminal.height - 10
			if listHeight < 3 {
				listHeight = 3
			}
		}
		m.list.SetSize(m.terminal.width, listHeight)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case internal.DebugMsg:
		m.debug = append(m.debug, string(msg))
		m.trimLog()
		return m, listenForDebug(m.debugCh)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) trimLog() {
	if len(m.debug) > 100 {
		m.debug = m.debug[len(m.debug)-100:]
	}
}
