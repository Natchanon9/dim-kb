---
title: "TCP Server"
description: "Raw TCP with net.Listen and net.Dial"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`net.Listen("tcp", addr)` gives a listener; each `Accept` returns a `net.Conn`, which is an `io.Reader` and `io.Writer`.

## Echo server

```go title="echo.go"
ln, err := net.Listen("tcp", ":9000")
if err != nil {
	log.Fatal(err)
}
for {
	conn, err := ln.Accept()
	if err != nil {
		continue
	}
	go func() { // one goroutine per connection
		defer conn.Close()
		line, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Fprintf(conn, "echo: %s", line)
	}()
}
// client: net.Dial("tcp", "localhost:9000"), send "hi\n" → "echo: hi"
```

TCP is a byte stream, not messages: you define framing yourself (newline-delimited, length prefix).

:::caution[Pitfalls]
- Set deadlines (`conn.SetDeadline`) so idle or slow clients don't hold goroutines forever.
:::
