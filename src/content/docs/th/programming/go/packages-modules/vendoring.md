---
title: "Vendoring"
description: "copy dependency เข้ามาใน repo ด้วย go mod vendor"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`go mod vendor` copy ทุก dependency มาไว้ในโฟลเดอร์ `vendor/` เมื่อมี `vendor/` และตรงกับ `go.mod` คำสั่ง go จะ build จากโฟลเดอร์นี้แทน module cache

## คำสั่ง

```sh
go mod vendor        # สร้างหรืออัปเดต vendor/
go build ./...       # ใช้ vendor/ ให้อัตโนมัติ (Go 1.14+)
go build -mod=mod    # ไม่สน vendor/ ใช้ module cache
```

## เมื่อไหร่ควร vendor

- build ต้องทำงานได้โดยไม่มี network (CI ที่ตัดขาดจากภายนอก)
- อยากให้เห็นโค้ดของ dependency ใน code review

โปรเจกต์ส่วนใหญ่ไม่ vendor เพราะ module proxy กับ `go.sum` ทำให้ build ซ้ำได้เหมือนเดิมอยู่แล้ว

:::caution[ข้อควรระวัง]
- แก้ `go.mod` แล้วต้องรัน `go mod vendor` ใหม่ ไม่อย่างนั้น build พังด้วย "inconsistent vendoring"
:::
