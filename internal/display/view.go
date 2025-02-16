package display

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	statusBar := "Shelf - Custom keybindings in one place"

	b.WriteString(strings.Repeat("─", m.terminal.width) + "\n")
	b.WriteString(centerText(statusBar, m.terminal.width) + "\n")
	b.WriteString(strings.Repeat("─", m.terminal.width) + "\n\n")

	m.paintKeybindings(&b)

	debugLines := m.paintDebugConsole(&b)

	m.fillRestOfHeightWithBlank(&b, debugLines)

	return b.String()
}

func (m Model) paintKeybindings(b *strings.Builder) {
	b.WriteString("Keybindings:\n\n")
	for i, kb := range m.keybindings {
		cursor := " "
		if m.cursor == i {
			cursor = ">" // Mark current keybinding with a cursor.
		}

		keys := strings.Join(kb.Keys, "+")

		paddedNamespace := fmt.Sprintf("%-15s", kb.Namespace)

		lineInfo := fmt.Sprintf("Line: %d,", kb.Breadcrumbs.Line)
		if kb.Breadcrumbs.Line == 0 {
			lineInfo = ""
		}

		line := fmt.Sprintf("%s %s: %s (File: %s,%s %d ms)\n",
			cursor,
			paddedNamespace,
			keys,
			kb.Breadcrumbs.FileName,
			lineInfo,
			kb.Telemetry.Parse.Milliseconds(),
		)
		b.WriteString(line)
	}

	if len(m.keybindings) == 0 {
		b.WriteString("No keybindings found.\n")
	}
}

func (m Model) paintDebugConsole(b *strings.Builder) int {
	if m.debugEnabled {
		b.WriteString("\n" + strings.Repeat("─", m.terminal.width) + "\n")
		b.WriteString(centerText("DEBUG PANEL", m.terminal.width) + "\n")
		b.WriteString(strings.Repeat("─", m.terminal.width) + "\n")

		maxDebugLines := 15
		startIdx := 0
		if len(m.debug) > maxDebugLines {
			startIdx = len(m.debug) - maxDebugLines
		}
		for _, msg := range m.debug[startIdx:] {
			b.WriteString(msg)
			if !strings.HasSuffix(msg, "\n") {
				b.WriteString("\n")
			}
		}

		return maxDebugLines
	}

	return 0
}

func (m Model) fillRestOfHeightWithBlank(b *strings.Builder, reserveForDebug int) {
	linesUsed := 6 + len(m.keybindings) + reserveForDebug
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
