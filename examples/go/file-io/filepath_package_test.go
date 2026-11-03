package fileio

import (
	"fmt"
	"path/filepath"
)

func Example_filepath() {
	p := filepath.Join("data", "users", "..", "orders", "2024.csv") // also cleans ".."
	fmt.Println(filepath.ToSlash(p), filepath.Base(p), filepath.Ext(p), filepath.ToSlash(filepath.Dir(p)))
	ok, _ := filepath.Match("*.csv", "report.csv")
	fmt.Println(ok)
	// Output:
	// data/orders/2024.csv 2024.csv .csv data/orders
	// true
}
