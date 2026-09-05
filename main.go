package main

import (
	"embed"
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
)

//go:embed banner.txt
var bannerFS embed.FS

//go:embed ignore
var ignoreFS embed.FS

var version = "v0.1.0"

func printBanner() {
	text, _ := bannerFS.ReadFile("banner.txt")
	fmt.Print(strings.ReplaceAll(string(text), "$version", version))
}

func printUsage() {
	fmt.Println("Usage: goignore <lang>\n Help: -h")
}

func generate(lang string) error {
	f, err := ignoreFS.ReadFile(path.Join("ignore", fmt.Sprintf("%s.gitignore", lang)))
	if err != nil {
		return fmt.Errorf("There is no .gitignore for %s", lang)
	}
	if err := os.WriteFile(".gitignore", f, 0644); err != nil {
		return fmt.Errorf("Failed to write file %s", err)
	}
	return nil
}

func main() {
	if slices.Contains(os.Args, "-h") {
		printBanner()
		printUsage()
		fmt.Print("\nList of generators:\n")
		templ, _ := ignoreFS.ReadDir("ignore")
		for _, tmp := range templ {
			r := strings.NewReplacer(".gitignore", "")
			tname := r.Replace(tmp.Name())
			parts := strings.Split(tname, ".")
			fmt.Println("-", parts[0])
		}
		return
	}
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	fname := os.Args[1]
	msg := "Created"
	if _, err := os.Stat(".gitignore"); err == nil {
		fmt.Print("File already exist, will be replaced? (y/n): ")
		var y string
		for {
			n, err := fmt.Scan(&y)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Err: %s\n", err)
				os.Exit(1)
			}
			if n != 1 {
				fmt.Printf("Please enter (y/n)")
				continue
			}
			if y != "y" {
				return
			}
			msg = "Replaced"
			break
		}
	}
	if err := generate(fname); err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s .gitignore file for %s 🎉\n", msg, fname)
}
