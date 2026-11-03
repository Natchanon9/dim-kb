package jsonencoding

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func Example_streaming() {
	dec := json.NewDecoder(strings.NewReader(`{"id":1} {"id":2}`))
	for {
		var e struct{ ID int }
		err := dec.Decode(&e)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(e.ID)
	}

	enc := json.NewEncoder(os.Stdout) // writes straight to any io.Writer
	enc.Encode(map[string]string{"status": "ok"})
	// Output:
	// 1
	// 2
	// {"status":"ok"}
}
