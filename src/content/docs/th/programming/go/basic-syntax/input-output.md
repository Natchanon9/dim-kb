---
title: "Input และ Output"
description: "พิมพ์ด้วย format verb ของ fmt และอ่าน input ด้วย fmt.Scan"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

package `fmt` ใช้ทั้งพิมพ์และอ่านค่า รู้ format verb ไม่กี่ตัวก็อ่าน debug output ได้หมด

## การพิมพ์

```go title="print.go"
u := user{"Ann", 30}
fmt.Printf("%v | %+v | %T\n", u, u, u)
// {Ann 30} | {Name:Ann Age:30} | main.user
fmt.Printf("[%6.2f] [%-4s] %q\n", 3.14159, "go", "hi")
// [  3.14] [go  ] "hi"
```

| Verb | พิมพ์ |
| --- | --- |
| `%v` / `%+v` / `%#v` | ค่า / พร้อมชื่อ field / แบบ syntax ของ Go |
| `%T` | type |
| `%d` `%f` `%s` `%q` `%t` | int, float, string, string ในเครื่องหมายคำพูด, bool |
| `%6.2f` / `%-4s` | ความกว้างและทศนิยม / ชิดซ้าย |

`Println` เว้นวรรคและขึ้นบรรทัดใหม่ให้ `Printf` จัด format `Sprintf` คืนค่าเป็น string และ `Fprintf` เขียนลง `io.Writer` อะไรก็ได้

## อ่าน input

```go title="scan.go"
var name string
var age int
fmt.Scan(&name, &age) // อ่าน "Bob 25" จาก stdin
```

ถ้าจะอ่านทั้งบรรทัด ใช้ `bufio.Scanner` กับ `os.Stdin`

:::caution[ข้อควรระวัง]
- `Scan` ต้องรับ pointer (`&name`) ส่งค่าไปตรงๆ ไม่มีประโยชน์
- เช็ค error ของ `Scan` ด้วย input ผิดรูปแบบจะไม่เปลี่ยนค่าตัวแปร
:::
