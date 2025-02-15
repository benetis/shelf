package loader

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const maxFiles = 100
const maxFileSize = 1 << 20 // 1 MB

type File struct {
	Contents []byte
	Name     string
	Path     string
}

func LoadFolder(folderPath string, debug bool) []File {
	folder := replaceTilde(folderPath)

	if debug {
		log.Print("Loading folder at", folder)
		log.Println("...")
	}

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if debug {
			log.Printf("Folder %q does not exist. Skipping...", folder)
		}
		return []File{}
	}

	var loadedFiles []File

	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == ".git" {
			return fs.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		if len(loadedFiles) >= maxFiles {
			return fmt.Errorf("too many files, max is %d", maxFiles)
		}

		file, err := loadFile(path)
		if err != nil {
			return fmt.Errorf("cannot load file %q: %w", path, err)
		}
		loadedFiles = append(loadedFiles, file)
		return nil
	})
	if err != nil {
		panic(fmt.Errorf("cannot walk folder: %w", err))
	}

	return loadedFiles
}

func loadFile(path string) (File, error) {
	f, err := os.Open(path)
	if err != nil {
		return File{}, fmt.Errorf("cannot open file: %w", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return File{}, fmt.Errorf("cannot get file info: %w", err)
	}

	if info.Size() > maxFileSize {
		return File{}, fmt.Errorf("file size (%d bytes) exceeds maximum allowed (%d bytes)", info.Size(), maxFileSize)
	}

	contents, err := io.ReadAll(f)
	if err != nil {
		return File{}, fmt.Errorf("cannot read file: %w", err)
	}

	return File{
		Contents: contents,
		Name:     info.Name(),
		Path:     path,
	}, nil
}

func replaceTilde(path string) string {
	if path == "~" {
		return os.Getenv("HOME")
	}
	if len(path) > 1 && path[:2] == "~/" {
		return filepath.Join(os.Getenv("HOME"), path[2:])
	}
	return path
}
