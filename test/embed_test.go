package golang_embed

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed golang_logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("golang_logo_new.png", logo, fs.ModePerm)

	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt

var multiFiles embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := multiFiles.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := multiFiles.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := multiFiles.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var files embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, err := files.ReadDir("files")

	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, err := os.ReadFile("files/" + entry.Name())

			if err != nil {
				panic(err)
			}

			fmt.Println("Content:", string(content))
		}
	}
}
