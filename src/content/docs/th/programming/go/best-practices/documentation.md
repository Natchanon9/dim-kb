---
title: "Documentation"
description: "doc comment, example และ pkg.go.dev"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

documentation ของ Go คือ comment ธรรมดา comment ที่อยู่เหนือการประกาศพอดีจะกลายเป็น doc ใน `go doc`, editor และ pkg.go.dev

## Doc comment

```go title="greet.go"
// Greet returns a greeting for name.
//
// It is safe for concurrent use. See also [Greeter.Greet].
func Greet(name string) string { return "Hello, " + name }
```

- ขึ้นต้นด้วยชื่อที่กำลังอธิบาย เขียนเป็นประโยคเต็ม
- `[Name]` ลิงก์ไปยังชื่ออื่น (Go 1.19+)
- comment `// Package foo ...` ในไฟล์ใดไฟล์หนึ่ง (มักเป็น `doc.go`) คือ doc ของ package

## Example คือ doc ที่ถูก test

```go title="greet_test.go"
func ExampleGreet() {
	fmt.Println(Greet("docs"))
	// Output: Hello, docs
}
```

`go test` รันและเช็ค output ให้ ส่วน pkg.go.dev แสดงไว้ใต้ `Greet`

:::caution[ข้อควรระวัง]
- comment ที่มีบรรทัดว่างคั่นกับการประกาศไม่นับเป็น doc comment
:::
