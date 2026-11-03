---
title: "Go Modules"
description: "go.mod, go.sum และคำสั่งจัดการ dependency"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

module คือกลุ่ม package ที่มี `go.mod` อยู่ที่ root ไฟล์นี้ระบุ module path, เวอร์ชัน Go และ dependency

## go.mod

```text title="go.mod"
module github.com/me/myapp

go 1.22

require github.com/gofiber/fiber/v2 v2.52.0
```

`go.sum` เก็บ checksum ของทุก dependency ทำให้ build ซ้ำได้เหมือนเดิมและรู้ถ้ามีใครแก้ไฟล์ ให้ commit ทั้งสองไฟล์

## คำสั่ง

```sh
go mod init github.com/me/myapp   # สร้าง go.mod
go get github.com/gofiber/fiber/v2 # เพิ่ม dependency
go mod tidy                        # เพิ่มที่ขาด ลบที่ไม่ใช้
go list -m all                     # ดู dependency ทั้งหมด
```

ไฟล์ที่โหลดจะไปอยู่ใน module cache (`$GOPATH/pkg/mod`) ผ่าน proxy `proxy.golang.org`

:::caution[ข้อควรระวัง]
- รัน `go mod tidy` ก่อน commit CI มักพังเพราะ `go.mod` ไม่ tidy
- repo ส่วนตัวต้องตั้ง `GOPRIVATE=github.com/yourorg/*` เพื่อไม่ให้วิ่งไป proxy สาธารณะ
:::
