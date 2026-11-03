---
title: "Constructor Functions"
description: "function NewXxx ที่ตรวจ input และตั้งค่าเริ่มต้น"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Go ไม่มี constructor ตามธรรมเนียมจะใช้ function ชื่อ `NewT` สร้าง `T` (ส่วนใหญ่เป็น `*T`) ตรวจ input และเติมค่าเริ่มต้น

## ตัวอย่าง

```go title="server.go"
type Server struct {
	addr    string
	timeout time.Duration
}

func NewServer(addr string) (*Server, error) {
	if addr == "" {
		return nil, errors.New("addr required")
	}
	return &Server{addr: addr, timeout: 30 * time.Second}, nil
}
```

## เมื่อไหร่ต้องมี

- มี field บังคับหรือต้อง validate
- ค่าเริ่มต้นไม่ใช่ zero value
- มี state ภายในที่ต้อง initialize (map, channel)

ถ้า zero value ใช้ได้อยู่แล้ว ไม่ต้องมี constructor

:::caution[ข้อควรระวัง]
- มี setting ที่เลือกได้เยอะ → ใช้ struct ของ option หรือ functional options ไม่ใช่ parameter สิบตัว
- constructor ที่ล้มเหลวได้ควรคืน `error` ไม่ใช่ panic
:::
