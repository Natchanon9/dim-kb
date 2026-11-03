package fileio

import (
	"fmt"
	"io/fs"
	"testing/fstest"
)

func Example_fs() {
	var fsys fs.FS = fstest.MapFS{ // in-memory fs.FS, handy for tests
		"a.txt":     {Data: []byte("A")},
		"dir/b.txt": {Data: []byte("B")},
	}
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			fmt.Println(p)
		}
		return nil
	})
	data, _ := fs.ReadFile(fsys, "dir/b.txt")
	fmt.Println(string(data))
	// Output:
	// a.txt
	// dir/b.txt
	// B
}
