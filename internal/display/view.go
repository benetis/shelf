package display

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	statusBar := "Shelf - A terminal keybinding cheat sheet"

	b.WriteString(strings.Repeat("─", m.terminal.width) + "\n")
	b.WriteString(centerText(statusBar, m.terminal.width) + "\n")
	b.WriteString(strings.Repeat("─", m.terminal.width) + "\n\n")

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

	m.fillRestOfHeightWithBlank(&b)

	return b.String()
}

func (m Model) fillRestOfHeightWithBlank(b *strings.Builder) {
	linesUsed := 6 + len(m.keybindings)
	for i := linesUsed; i < m.terminal.height; i++ {
		b.WriteString("\n")
	}
}

func centerText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	spaces := (width - len(text)) / 2
	return fmt.Sprintf("%s%s", strings.Repeat(" ", spaces), text)
}
