package display

import (
	"fmt"
	"github.com/benetis/shelf/internal"
	"strings"
)

type keybindingItem struct {
	kb internal.Keybinding
}

func (i keybindingItem) Title() string {
	return fmt.Sprintf("%s: %s", i.kb.Namespace, strings.Join(i.kb.Keys, "+"))
}

func (i keybindingItem) Description() string {
	lineInfo := ""
	if i.kb.Breadcrumbs.Line != 0 {
		lineInfo = fmt.Sprintf("Line: %d,", i.kb.Breadcrumbs.Line)
	}
	return fmt.Sprintf("(File: %s, %s %d ms)",
		i.kb.Breadcrumbs.FileName,
		lineInfo,
		i.kb.Telemetry.Parse.Milliseconds(),
	)
}

func (i keybindingItem) FilterValue() string {
	return i.kb.Namespace
}

func (m Model) View() string {
	var b strings.Builder

	border := strings.Repeat("─", m.terminal.width)

	b.WriteString(m.list.View())

	if m.debugEnabled {
		b.WriteString("\n" + border + "\n")
		b.WriteString(centerText("DEBUG PANEL", m.terminal.width) + "\n")
		b.WriteString(border + "\n")
		for _, msg := range m.debug {
			b.WriteString(msg)
			if !strings.HasSuffix(msg, "\n") {
				b.WriteString("\n")
			}
		}
	}

	viewLines := strings.Split(b.String(), "\n")
	for len(viewLines) < m.terminal.height {
		viewLines = append(viewLines, "")
	}
	return strings.Join(viewLines, "\n")
}

func centerText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	spaces := (width - len(text)) / 2
	return fmt.Sprintf("%s%s", strings.Repeat(" ", spaces), text)
}
