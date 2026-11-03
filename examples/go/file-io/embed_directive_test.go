package fileio

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed static/hello.txt
var hello string

//go:embed static
var staticFS embed.FS

func Example_embed() {
	fmt.Print(hello)
	entries, _ := fs.ReadDir(staticFS, "static")
	fmt.Println(len(entries), entries[0].Name())
	// Output:
	// hello from embed
	// 1 hello.txt
}
