---
title: "Cross-Compilation"
description: "build ให้ OS และ CPU อื่นด้วย GOOS และ GOARCH"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Go cross-compile ได้ทันที: ตั้ง `GOOS` กับ `GOARCH` แล้ว build โค้ดที่เป็น Go ล้วนไม่ต้องมี toolchain เพิ่ม

## คำสั่ง

```sh
GOOS=linux   GOARCH=amd64 go build -o app-linux .
GOOS=linux   GOARCH=arm64 go build -o app-arm64 .   # Graviton, Raspberry Pi
GOOS=darwin  GOARCH=arm64 go build -o app-mac .
GOOS=windows GOARCH=amd64 go build -o app.exe .
go tool dist list                                   # ทุกคู่ที่รองรับ
```

## Build สำหรับ production

```sh
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X main.version=1.4.2" -o app ./cmd/api
```

| Flag | ผล |
| --- | --- |
| `CGO_ENABLED=0` | binary แบบ static รันบน image `scratch`/distroless ได้ |
| `-trimpath` | ตัด path ในเครื่องออกจาก binary |
| `-ldflags="-s -w"` | ตัด debug info: binary เล็กลง |
| `-X main.version=...` | ตั้งค่าตัวแปร string ตอน build |

:::caution[ข้อควรระวัง]
- package ที่ใช้ cgo cross-compile ง่ายๆ แบบนี้ไม่ได้ ต้องมี C cross-compiler หรือใช้ตัวที่เป็น Go ล้วนแทน
:::
