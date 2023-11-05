package golang_embed

import (
	_ "embed"
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
