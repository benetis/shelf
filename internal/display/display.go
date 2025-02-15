package display

import (
	"fmt"
	"github.com/benetis/shelf/internal"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type model struct {
	keybindings []internal.Keybinding
	cursor      int
}

func InitialModel(keybindings []internal.Keybinding) model {
	return model{
		keybindings: keybindings,
		cursor:      0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString("Keybindings:\n\n")

	for i, kb := range m.keybindings {
		cursor := " "
		if m.cursor == i {
			cursor = ">" // Mark current keybinding with a cursor.
		}

		modStr := strings.Join(kb.Modifiers, "+")
		line := fmt.Sprintf("%s %s: %s (File: %s, Line: %d)\n",
			cursor,
			modStr,
			kb.Key,
			kb.Breadcrumbs.FileName,
			kb.Breadcrumbs.Line,
		)
		b.WriteString(line)
	}

	b.WriteString("\nPress q to quit.\n")
	return b.String()
}
