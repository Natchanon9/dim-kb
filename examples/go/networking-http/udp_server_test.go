package networkinghttp

import (
	"bytes"
	"fmt"
	"net"
)

func Example_udp() {
	pc, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer pc.Close()
	go func() {
		buf := make([]byte, 64)
		n, addr, err := pc.ReadFrom(buf) // one datagram, plus who sent it
		if err != nil {
			return
		}
		pc.WriteTo(bytes.ToUpper(buf[:n]), addr)
	}()

	conn, _ := net.Dial("udp", pc.LocalAddr().String())
	defer conn.Close()
	conn.Write([]byte("ping"))
	buf := make([]byte, 64)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
	// Output: PING
}
