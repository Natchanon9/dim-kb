---
title: "ติดตั้ง Go"
description: "ติดตั้ง Go เช็คเวอร์ชัน และรู้ความต่างของ GOROOT กับ GOPATH"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

โหลด Go จาก [go.dev/dl](https://go.dev/dl) หรือใช้ package manager แล้วเช็คว่าคำสั่ง `go` อยู่ใน `PATH`

## ติดตั้งบน Linux

```sh
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.xx.x.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin   # ใส่ไว้ใน ~/.profile หรือ ~/.zshrc
go version
```

macOS และ Windows มี installer ให้ในหน้าเดียวกัน

## Path ที่ Go ใช้

| ตัวแปร | ค่าเริ่มต้น | เก็บอะไร |
| --- | --- | --- |
| `GOROOT` | ที่ติดตั้ง (`/usr/local/go`) | ตัว Go toolchain |
| `GOPATH` | `~/go` | module cache (`pkg/mod`) และ binary จาก `go install` (`bin`) |

เพิ่ม `$(go env GOPATH)/bin` เข้า `PATH` ด้วย เครื่องมือที่ติดตั้งด้วย `go install` จะได้รันได้ ส่วน `go env` แสดงค่าทั้งหมด

## เวอร์ชันของ toolchain

`go.mod` ของแต่ละ module ระบุเวอร์ชัน Go ขั้นต่ำ (`go 1.22`) ตั้งแต่ Go 1.21 ถ้า Go ในเครื่องเก่ากว่า คำสั่ง `go` จะโหลด toolchain ที่ต้องใช้มาให้เอง

:::caution[ข้อควรระวัง]
- อย่าตั้ง `GOROOT` เองถ้าไม่ได้มีหลายเวอร์ชัน ตั้งผิดแล้ว toolchain พัง
- พอใช้ modules แล้ว โค้ดวางที่ไหนก็ได้ ไม่ต้องอยู่ใน `GOPATH`
:::
