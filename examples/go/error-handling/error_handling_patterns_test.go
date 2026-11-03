package errorhandling

import (
	"errors"
	"fmt"
)

func readFile(path string) (string, error) {
	if path == "" {
		return "", errors.New("empty path")
	}
	return "data", nil
}

func loadConfig(path string) (string, error) {
	data, err := readFile(path)
	if err != nil {
		return "", fmt.Errorf("load config: %w", err) // add context, return early
	}
	return data, nil
}

func Example_patterns() {
	_, err := loadConfig("")
	fmt.Println(err)
	// Output: load config: empty path
}
