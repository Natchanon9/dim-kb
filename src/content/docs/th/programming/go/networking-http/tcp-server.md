---
title: "TCP Server"
description: "TCP ดิบด้วย net.Listen และ net.Dial"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`net.Listen("tcp", addr)` ได้ listener มา แต่ละ `Accept` คืน `net.Conn` ซึ่งเป็นทั้ง `io.Reader` และ `io.Writer`

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
	go func() { // หนึ่ง goroutine ต่อหนึ่ง connection
		defer conn.Close()
		line, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Fprintf(conn, "echo: %s", line)
	}()
}
// client: net.Dial("tcp", "localhost:9000") ส่ง "hi\n" → "echo: hi"
```

TCP คือ stream ของ byte ไม่ใช่ข้อความเป็นก้อน เราต้องกำหนดการแบ่งข้อความเอง (คั่นด้วยขึ้นบรรทัดใหม่, ใส่ความยาวนำหน้า)

:::caution[ข้อควรระวัง]
- ตั้ง deadline (`conn.SetDeadline`) client ที่เงียบหรือช้าจะได้ไม่ถือ goroutine ไว้ตลอดไป
:::
