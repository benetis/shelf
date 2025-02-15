package internal

type KeyBinding struct {
	Modifiers   []string
	Key         string
	Breadcrumbs Breadcrumbs
}

type Breadcrumbs struct {
	FileName string
	Line     int
}
