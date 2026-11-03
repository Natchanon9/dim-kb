---
title: "Custom Error Types"
description: "struct ของ error ที่พก field เพิ่มเติม"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

ถ้าผู้เรียกต้องการรายละเอียด (field ไหน status code อะไร) ให้คืน struct ที่ implement `error`

## ตัวอย่าง

```go title="validation.go"
type ValidationError struct {
	Field, Reason string
}

func (e *ValidationError) Error() string { return e.Field + ": " + e.Reason }

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return &ValidationError{Field: "email", Reason: "missing @"}
	}
	return nil
}
// validateEmail("bob") → email: missing @
```

ผู้เรียกดึงออกมาด้วย `errors.As` ถ้า type ของเรา wrap error อื่นอยู่ ให้เพิ่ม `Unwrap() error` เพื่อให้ `Is`/`As` มองเข้าไปข้างในได้

:::caution[ข้อควรระวัง]
- ตอนสำเร็จให้คืน `nil` ไม่ใช่ `*ValidationError` ที่เป็น nil typed nil ทำให้ `err != nil` เป็นจริง
- ใช้ pointer receiver และคืน `&T{}` เป้าหมายของ `errors.As` จะได้สม่ำเสมอ
:::
