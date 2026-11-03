package fileio

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Example_tempFiles() {
	f, _ := os.CreateTemp("", "upload-*.txt") // * is replaced with a random string
	defer os.Remove(f.Name())
	f.WriteString("tmp")
	f.Close()
	fmt.Println(strings.HasPrefix(filepath.Base(f.Name()), "upload-"), filepath.Ext(f.Name()))
	// Output: true .txt
}

func TestTempDir(t *testing.T) {
	dir := t.TempDir() // removed automatically when the test ends
	if err := os.WriteFile(filepath.Join(dir, "x"), []byte("1"), 0o600); err != nil {
		t.Fatal(err)
	}
}
