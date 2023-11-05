package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed golang_logo.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("generated_logo.png", logo, fs.ModePerm)
	if err != nil {
		panic(nil)
	}

	entries, _ := path.ReadDir("files")

	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			content, _ := path.ReadFile("files/" + entry.Name())

			fmt.Println(string(content))
		}
	}
}
