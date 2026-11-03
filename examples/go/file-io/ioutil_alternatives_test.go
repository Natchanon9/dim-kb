package fileio

import (
	"fmt"
	"os"
	"path/filepath"
)

func Example_ioutilAlternatives() {
	dir, _ := os.MkdirTemp("", "list") // was ioutil.TempDir
	defer os.RemoveAll(dir)
	for _, name := range []string{"b.txt", "a.txt"} {
		os.WriteFile(filepath.Join(dir, name), nil, 0o644) // was ioutil.WriteFile
	}
	entries, _ := os.ReadDir(dir) // was ioutil.ReadDir; sorted by name
	for _, e := range entries {
		fmt.Println(e.Name())
	}
	// Output:
	// a.txt
	// b.txt
}
