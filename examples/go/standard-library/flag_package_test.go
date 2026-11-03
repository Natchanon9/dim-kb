package standardlibrary

import (
	"flag"
	"fmt"
)

func Example_flag() {
	fs := flag.NewFlagSet("server", flag.ContinueOnError) // flag.Int etc. use os.Args in main
	port := fs.Int("port", 8080, "listen port")
	debug := fs.Bool("debug", false, "enable debug logs")
	fs.Parse([]string{"-port=9000", "-debug", "extra"})
	fmt.Println(*port, *debug, fs.Args())
	// Output: 9000 true [extra]
}
