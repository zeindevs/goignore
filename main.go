package main

import (
	"embed"
	"fmt"
	"os"
	"path"
)

//go:embed ignore
var ignoreFS embed.FS

func printUsage() {
	msg := `Usage: goignore [lang]

Lang available: go,node,python,c,cpp

Example: goignore go`
	fmt.Println(msg)
}

func generate(lang string) error {
	f, err := ignoreFS.ReadFile(path.Join("ignore", fmt.Sprintf("%s.gitignore", lang)))
	if err != nil {
		return fmt.Errorf("There is no .gitignore for %s", lang)
	}

	if err := os.WriteFile(".gitignore", f, 0644); err != nil {
		return fmt.Errorf("Failed to write file %s", err.Error())
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		return
	}

	if err := generate(os.Args[1]); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Created .gitignore file for %s 🎉\n", os.Args[1])
}
