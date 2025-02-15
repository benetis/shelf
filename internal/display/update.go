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
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.keybindings)-1 {
				m.cursor++
			}
		}
	case internal.DebugMsg:
		m.debug = append(m.debug, string(msg))

		m.trimLog()

		return m, listenForDebug(m.debugCh)
	}
	return m, nil
}

func (m Model) trimLog() {
	if len(m.debug) > 100 {
		m.debug = m.debug[len(m.debug)-100:]
	}
}
