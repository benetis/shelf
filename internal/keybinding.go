package internal

import "time"

type Keybinding struct {
	Keys        []string
	Namespace   string
	Metadata    string
	Breadcrumbs Breadcrumbs
	Telemetry   Telemetry
}

type Breadcrumbs struct {
	FileName string
	Line     int
}

type Telemetry struct {
	Parse time.Duration
}
