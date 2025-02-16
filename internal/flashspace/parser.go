package flashspace

import (
	"encoding/json"
	"github.com/benetis/shelf/internal"
	"github.com/benetis/shelf/internal/loader"
	"log"
	"time"
)

type flashspace struct {
	Profiles []profile `json:"profiles"`
}

type profile struct {
	Name       string      `json:"name"`
	Workspaces []workspace `json:"workspaces"`
}

type workspace struct {
	Name     string `json:"name"`
	Shortcut string `json:"shortcut"`
}

func Parse(filePath string, debug bool) []internal.Keybinding {
	startLoad := time.Now()
	var keybindings []internal.Keybinding

	file, err := loader.LoadFile(filePath)
	if err != nil {
		log.Printf("Error loading flashspace file: %s\n", err)
		return keybindings
	}
	loadDuration := time.Since(startLoad)

	if debug {
		log.Printf("Loaded flashspace file in %s\n", loadDuration)
	}

	var f flashspace

	err = json.Unmarshal(file.Contents, &f)
	if err != nil {
		log.Printf("Error parsing flashspace file: %s\n", err)
		return keybindings
	}

	endParse := time.Since(startLoad)

	for _, p := range f.Profiles {
		for _, w := range p.Workspaces {
			keybindings = append(keybindings, internal.Keybinding{
				Modifiers: []string{},
				Key:       w.Shortcut,
				Breadcrumbs: internal.Breadcrumbs{
					FileName: file.Name,
					Line:     0,
				},
				Telemetry: internal.Telemetry{
					Parse: endParse,
				},
			})
		}
	}

	if debug {
		log.Printf("Found %d flashspace keybindings\n", len(keybindings))
	}

	return keybindings
}
