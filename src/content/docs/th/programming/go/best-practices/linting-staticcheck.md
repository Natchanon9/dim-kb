---
title: "Linting และ staticcheck"
description: "go vet, staticcheck และ golangci-lint ใน CI"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

linter จับ bug ที่ compiler ยอมให้ผ่าน รันทั้งในเครื่องและใน CI

## เครื่องมือ

| เครื่องมือ | จับอะไร |
| --- | --- |
| `go vet` | Printf ผิด, copy lock, โค้ดที่ไม่มีทางถึง (มากับ Go) |
| `staticcheck` | bug, API ที่ deprecated, จุดที่เขียนให้ง่ายลงได้ |
| `golangci-lint` | รัน linter หลายตัว (รวม staticcheck, errcheck, gosec) ในรอบเดียวที่เร็ว |

```sh
go vet ./...
go install honnef.co/go/tools/cmd/staticcheck@latest && staticcheck ./...
golangci-lint run
```

## Config ขั้นต่ำ

```yaml title=".golangci.yml"
version: "2"
linters:
  enable:
    - errcheck
    - staticcheck
    - gosec
    - revive
```

:::caution[ข้อควรระวัง]
- เปิด linter ทุกตัวจะได้ noise จนคนเลิกสนใจ เริ่มจากน้อยๆ แล้วเพิ่มกฎที่จับ bug จริงได้
:::
