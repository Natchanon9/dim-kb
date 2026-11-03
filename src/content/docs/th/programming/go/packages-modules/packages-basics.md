---
title: "พื้นฐาน Package"
description: "หนึ่งโฟลเดอร์ หนึ่ง package และตัวพิมพ์ใหญ่คือ export"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

package คือโฟลเดอร์ของไฟล์ `.go` ที่ใช้ชื่อ `package` เดียวกัน เป็นหน่วยหลักของการจัดโค้ดและการมองเห็นใน Go

## โครงสร้าง

```text
myapp/
├─ go.mod            module github.com/me/myapp
├─ main.go           package main
└─ greet/
   └─ greet.go       package greet
```

```go title="greet/greet.go"
package greet

// Hello เป็น exported: package อื่นเรียก greet.Hello ได้
func Hello(name string) string { return "Hello, " + name }

// suffix เป็น unexported: ใช้ได้เฉพาะโค้ดใน package greet
func suffix() string { return "!" }
```

```go title="main.go"
import "github.com/me/myapp/greet" // module path + โฟลเดอร์

fmt.Println(greet.Hello("gopher")) // Hello, gopher
```

## กติกา

- ทุกไฟล์ในโฟลเดอร์เป็น package เดียวกัน (ยกเว้น package `_test`)
- ชื่อ package ส่วนใหญ่ตรงกับชื่อโฟลเดอร์: สั้น ตัวพิมพ์เล็ก ไม่มี underscore
- การมองเห็นคิดเป็นราย **package** ไม่ใช่รายไฟล์

:::caution[ข้อควรระวัง]
- import วนกัน (`a` import `b` แล้ว `b` import `a`) compile ไม่ผ่าน ให้ย้ายโค้ดที่ใช้ร่วมไปไว้ package ที่สาม
- เลี่ยงชื่อรวมมิตรอย่าง `utils`, `common`, `helpers`
:::
