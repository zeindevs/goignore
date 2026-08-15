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

	fname := os.Args[1]

	msg := "Created"
	if _, err := os.Stat(".gitignore"); err == nil {
		var y string
		fmt.Print("File already exist, will be replaced? (y/n): ")
		n, err := fmt.Scan(&y)
		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
			return
		}
		if n != 1 {
			fmt.Printf("Please enter (y/n)")
			return
		}
		if y == "n" {
			return
		}
		msg = "Replaced"
	}

	if err := generate(fname); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s .gitignore file for %s 🎉\n", msg, fname)
}
