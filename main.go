package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var version = "dev"

//go:embed ignore
var ignoreFS embed.FS

func main() {
	var showVer bool

	flag.BoolVar(&showVer, "v", false, "show version")
	flag.Usage = func() {
		fmt.Print(fmt.Sprintf(`
                _
   ____ _____  (_)___ _____  ____  ________
  / __ '/ __ \/ / __ '/ __ \/ __ \/ ___/ _ \
 / /_/ / /_/ / / /_/ / / / / /_/ / /  /  __/
 \__, /\____/_/\__, /_/ /_/\____/_/   \___/
/____/        /____/

Usage:
  goignore go
  goignore cpp

Generators:
%s
Options:
`, generators()))
		flag.PrintDefaults()
	}
	flag.Parse()

	if showVer {
		fmt.Println("version", version)
		return
	}

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	lang := flag.Arg(0)

	templ, err := readIgnore(lang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	if err := writeIgnore(lang, templ); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}

func generators() string {
	var sb strings.Builder

	templ, _ := ignoreFS.ReadDir("ignore")

	for _, tmp := range templ {
		r := strings.NewReplacer(".gitignore", "")

		tname := r.Replace(tmp.Name())
		parts := strings.Split(tname, ".")

		sb.WriteString(fmt.Sprintf("  %s\n", parts[0]))
	}

	return sb.String()
}

func readIgnore(lang string) ([]byte, error) {
	templ, err := ignoreFS.ReadFile(path.Join("ignore", fmt.Sprintf("%s.gitignore", lang)))
	if err != nil {
		return nil, fmt.Errorf("There is no .gitignore for %s", lang)
	}
	return templ, nil
}

func writeIgnore(lang string, templ []byte) error {
	msg := "Created"

	if _, err := os.Stat(".gitignore"); err == nil {
		var y string

		fmt.Print("File already exist, will be replaced? (y/n): ")

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
				return nil
			}

			break
		}

		msg = "Replaced"
	}

	if err := os.WriteFile(".gitignore", templ, 0644); err != nil {
		return fmt.Errorf("Failed to write file %s", err)
	}

	fmt.Printf("%s .gitignore file for %s 🎉\n", msg, lang)

	return nil
}
