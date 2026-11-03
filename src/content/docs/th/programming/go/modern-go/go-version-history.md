---
title: "ประวัติเวอร์ชัน Go"
description: "แต่ละ release ล่าสุดเพิ่มอะไรที่ได้ใช้จริง"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Go ออก release ทุกกุมภาพันธ์และสิงหาคม ทุกตัวยังเข้ากันได้กับโค้ด Go 1.0 บรรทัด `go` ใน `go.mod` เป็นตัวตัดสินว่า module นั้นใช้ฟีเจอร์ของภาษาได้ถึงแค่ไหน

## Release ที่เปลี่ยนโค้ดประจำวัน

| เวอร์ชัน | ไฮไลต์ |
| --- | --- |
| 1.18 (2022) | generics, fuzzing, workspaces |
| 1.19 | atomic แบบมี type (`atomic.Int64`), syntax ของ doc comment |
| 1.20 | `errors.Join`, `%w` หลายตัว, `context.WithCancelCause` |
| 1.21 | `min`/`max`/`clear`, `slices`, `maps`, `cmp`, `log/slog`, โหลด toolchain อัตโนมัติ |
| 1.22 (2024) | loop var แยกต่อรอบ, `range` บน int, pattern ของ ServeMux, `math/rand/v2` |
| 1.23 | iterator แบบ range-over-func, `iter`, `unique` |
| 1.24 (2025) | generic type alias, `omitzero`, `os.Root`, `testing.B.Loop`, map แบบ Swiss table |
| 1.25 | `GOMAXPROCS` ที่รู้จัก container, `sync.WaitGroup.Go`, `testing/synctest` |

## ตามให้ทัน

- อ่าน release notes ที่ go.dev/doc/devel/release
- รัน modernizer ของ `go fix` / `gopls` เพื่อปรับโค้ดเป็นสำนวนใหม่
- เปลี่ยนบรรทัด `go` อย่างตั้งใจ เพราะเปลี่ยนความหมายของโค้ดได้ (เช่น loop variable)

:::caution[ข้อควรระวัง]
- ใช้ function ของ standard library ที่ใหม่กว่าบรรทัด `go` อนุญาต `go vet` (stdversion) จะรายงาน
:::
