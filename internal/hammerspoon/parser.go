package hammerspoon

import (
	"github.com/benetis/shelf/internal"
	"github.com/benetis/shelf/internal/loader"
	"log"
	"regexp"
	"strings"
	"time"
)

func Parse(folderPath string, debug bool) []internal.Keybinding {
	startLoad := time.Now()
	var keybindings []internal.Keybinding

	files := loader.LoadFolder(folderPath, debug)
	loadDuration := time.Since(startLoad)

	if debug {
		log.Printf("Loaded %d files in %s\n", len(files), loadDuration)
	}

	re := regexp.MustCompile(
		`hs\.hotkey\.bind\(` + // Match 'hs.hotkey.bind('
			`\{([^}]+)\}` + // Capture the modifiers inside curly braces
			`,\s*` + // Match a comma followed by optional whitespace
			`"([^"]+)`, // Capture the key inside double quotes
	)

	for _, f := range files {
		startFileParse := time.Now()
		lines := strings.Split(string(f.Contents), "\n")
		for i, line := range lines {
			result := oneLine(line, re, f, i, debug)
			if result != nil {
				result.Telemetry = internal.Telemetry{
					Parse: time.Since(startFileParse),
				}

				keybindings = append(keybindings, *result)
			}
		}
	}

	if debug {
		log.Printf("Found %d keybindings\n", len(keybindings))
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
				log.Printf("%s:%d: unmatched line: %s\n", f.Path, i+1, line)
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
