package main

import (
	"embed"
	"fmt"
	"os"
	"path"
)

func usage() {
	msg := `
Usage: goignore [PROJECT TYPE]
Types Available: go,node,python
Example: goignore go
  `
	fmt.Println(msg)
}

//go:embed ignore
var IgnoreFS embed.FS

func generate(lang string) error {
	f, err := IgnoreFS.ReadFile(path.Join("ignore", fmt.Sprintf("%s.txt", lang)))
	if err != nil {
		return fmt.Errorf("There is no gitignore for %s", lang)
	}
	fmt.Printf("Created .gitignore file for %s 🎉\n", lang)
	if err := os.WriteFile(".gitignore", f, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write file %s", err.Error())
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	if err := generate(os.Args[1]); err != nil {
		fmt.Println(err.Error())
	}
}
