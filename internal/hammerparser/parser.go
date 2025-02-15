package hammerparser

import (
	"fmt"
)

type KeyBinding struct {
	Modifiers []string
	Key       string
}

func Parse(folderPath string) {
	fmt.Println("Parsing...")

	files := loadFolder(folderPath)

	for _, file := range files {
		fmt.Println(file.Path)
	}

	//chunk, err := parse.Parse(loadFile(), "init.lua")
	//if err != nil {
	//	panic(err)
	//}
}
