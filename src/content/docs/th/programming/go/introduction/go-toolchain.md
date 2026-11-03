---
title: "Go Toolchain"
description: "คำสั่ง go ที่ใช้ทุกวัน"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

binary ตัวเดียวคือ `go` ทำได้ทั้ง build, test, format และจัดการ dependency ไม่ต้องมี build tool แยก

## คำสั่งที่ใช้ทุกวัน

| คำสั่ง | ทำอะไร |
| --- | --- |
| `go run .` | compile และรัน package ในโฟลเดอร์นี้ |
| `go build -o app .` | compile เป็น binary |
| `go test ./...` | รัน test ทั้งหมด (`-race`, `-cover`, `-run Name`) |
| `go fmt ./...` | format โค้ด (เหมือน `gofmt -w`) |
| `go vet ./...` | จับโค้ดน่าสงสัยที่ compiler ยอมให้ผ่าน |
| `go mod init <path>` | เริ่ม module |
| `go mod tidy` | เพิ่ม dependency ที่ขาด / ลบที่ไม่ใช้ |
| `go get pkg@v1.2.3` | เพิ่มหรือเปลี่ยนเวอร์ชัน dependency |
| `go install pkg@latest` | build เครื่องมือไปไว้ที่ `$GOPATH/bin` |
| `go env` | ดูค่าต่างๆ เช่น `GOPATH`, `GOOS` |
| `go doc fmt.Println` | ดู documentation |

## `./...`

`./...` แปลว่า "โฟลเดอร์นี้และทุก package ข้างใต้" คำสั่งส่วนใหญ่รับแบบนี้ได้

:::caution[ข้อควรระวัง]
- ใส่ `gofmt` กับ `go vet` ไว้ใน CI โค้ดที่ไม่ format เป็นเรื่องที่ทุกทีม Go ทักใน review
- `go get` ไม่ติดตั้ง binary แล้ว ให้ใช้ `go install pkg@version`
:::
