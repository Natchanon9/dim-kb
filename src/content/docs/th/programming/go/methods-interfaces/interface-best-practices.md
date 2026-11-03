---
title: "Interface Best Practices"
description: "interface เล็กๆ ประกาศโดยฝั่งที่ใช้"
sidebar:
  order: 11
status: done
lastReviewed: 2026-11-03
---

interface ที่ดีใน Go มีขนาดเล็ก และประกาศไว้ตรงที่ใช้ ไม่ใช่ข้างๆ ตัว implementation

## ประกาศ interface ที่ฝั่งผู้ใช้

```go title="notify.go"
// ประกาศโดยโค้ดที่ต้องใช้ และมีแค่ method ที่ใช้จริง
type UserStore interface {
	Email(id int) (string, error)
}

func Notify(s UserStore, id int) string {
	e, err := s.Email(id)
	if err != nil {
		return "no email"
	}
	return "sent to " + e
}

type fakeStore struct{}

func (fakeStore) Email(int) (string, error) { return "a@b.com", nil }

Notify(fakeStore{}, 1) // sent to a@b.com
```

## หลักง่ายๆ

- **รับ interface คืน struct**
- **ยิ่งเล็กยิ่งดี:** method เดียวดีที่สุด
- **อย่าสร้าง interface ก่อนจะมี implementation ตัวที่สอง (หรือตัวปลอมใน test) มาต้องการ**
- **ไม่ตั้งชื่อแบบ `IUserService`** ให้ตั้งตามพฤติกรรม

:::caution[ข้อควรระวัง]
- package ที่ export interface พร้อม implementation ตัวเดียวของมัน ส่วนใหญ่ต้องการแค่ struct
:::
