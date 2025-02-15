package hammerparser

import (
	"fmt"
	"github.com/benetis/shelf/internal"
	"github.com/benetis/shelf/internal/loader"
	"regexp"
	"strings"
)

func Parse(folderPath string, debug bool) []internal.Keybinding {
	var keybindings []internal.Keybinding

	files := loader.LoadFolder(folderPath)

	re := regexp.MustCompile(
		`hs\.hotkey\.bind\(` + // Match 'hs.hotkey.bind('
			`\{([^}]+)\}` + // Capture the modifiers inside curly braces
			`,\s*` + // Match a comma followed by optional whitespace
			`"([^"]+)`, // Capture the key inside double quotes
	)

	for _, f := range files {
		lines := strings.Split(string(f.Contents), "\n")
		for i, line := range lines {
			result := oneLine(line, re, f, i, debug)
			if result != nil {
				keybindings = append(keybindings, *result)
			}
		}
	}

	if debug {
		fmt.Printf("Found %d keybindings\n", len(keybindings))
	}

	return keybindings
}

func oneLine(line string, re *regexp.Regexp, f loader.File, i int, debug bool) *internal.Keybinding {
	if strings.Contains(line, "hs.hotkey.bind(") {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			modifiersStr := matches[1]
			key := matches[2]

			modifiers := parseModifiers(modifiersStr)

			binding := internal.Keybinding{
				Modifiers: modifiers,
				Key:       key,
				Breadcrumbs: internal.Breadcrumbs{
					FileName: f.Name,
					Line:     i + 1,
				},
			}

			return &binding
		} else {
			if debug {
				fmt.Printf("%s:%d: unmatched line: %s\n", f.Path, i+1, line)
			}
			return nil
		}
	}
	return nil
}

func parseModifiers(modifiersStr string) []string {
	parts := strings.Split(modifiersStr, ",")
	var modifiers []string
	for _, p := range parts {
		mod := strings.Trim(p, ` "'`)
		if mod != "" {
			modifiers = append(modifiers, mod)
		}
	}
	return modifiers
}
