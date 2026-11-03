package networkinghttp

import (
	"bufio"
	"fmt"
	"net"
)

func Example_tcp() {
	ln, err := net.Listen("tcp", "127.0.0.1:0") // :0 picks a free port
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	go func() {
		conn, err := ln.Accept() // a real server loops and handles each conn in a goroutine
		if err != nil {
			return
		}
		defer conn.Close()
		line, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Fprintf(conn, "echo: %s", line)
	}()

	conn, _ := net.Dial("tcp", ln.Addr().String())
	defer conn.Close()
	fmt.Fprint(conn, "hi\n")
	reply, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(reply)
	// Output: echo: hi
}
