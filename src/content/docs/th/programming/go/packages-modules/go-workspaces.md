---
title: "Go Workspaces"
description: "พัฒนาหลาย module ในเครื่องพร้อมกันด้วย go.work"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

workspace (Go 1.18+) ให้ build หลาย module บนดิสก์เหมือนเป็นก้อนเดียว โดยไม่ต้องใส่ `replace` ใน `go.mod`

## Setup

```sh
go work init ./api ./shared   # สร้าง go.work
go work use ./worker          # เพิ่ม module
```

```text title="go.work"
go 1.22

use (
	./api
	./shared
	./worker
)
```

ตอนนี้ `api` import โฟลเดอร์ `shared` ในเครื่องได้เลย แก้แล้วเห็นผลทันที

## ใช้เมื่อไหร่

แก้ library พร้อมกับ service ที่ใช้มันในเวลาเดียวกัน หรือ monorepo ที่มีหลาย module

:::caution[ข้อควรระวัง]
- หลายทีมไม่เอา `go.work` เข้า git เพื่อให้ CI build แต่ละ module ตาม `go.mod` จริง ตั้ง `GOWORK=off` ถ้าจะ test โดยไม่ใช้มัน
:::
