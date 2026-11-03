---
title: "ทางเลือกแทน ioutil"
description: "function ของ io/ioutil ย้ายไปอยู่ไหนใน Go 1.16"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`io/ioutil` เลิกใช้ตั้งแต่ Go 1.16 function ทั้งหมดย้ายไปอยู่ใน `io` และ `os` โดยทำงานเหมือนเดิม

## ตารางย้าย

| เดิม (`ioutil`) | ใหม่ |
| --- | --- |
| `ReadAll` | `io.ReadAll` |
| `ReadFile` / `WriteFile` | `os.ReadFile` / `os.WriteFile` |
| `ReadDir` | `os.ReadDir` (คืน `[]os.DirEntry` ถูกกว่า) |
| `TempDir` / `TempFile` | `os.MkdirTemp` / `os.CreateTemp` |
| `NopCloser` | `io.NopCloser` |
| `Discard` | `io.Discard` |

```go title="readdir.go"
dir, _ := os.MkdirTemp("", "list")
defer os.RemoveAll(dir)
entries, _ := os.ReadDir(dir) // เรียงตามชื่อ
for _, e := range entries {
	fmt.Println(e.Name())
}
```

:::caution[ข้อควรระวัง]
- `os.ReadDir` คืน `DirEntry` ไม่ใช่ `FileInfo` ถ้าต้องการขนาดหรือเวลาแก้ไขให้เรียก `e.Info()`
:::
