---
title: "Field แบบ Exported และ Unexported"
description: "field ตัวพิมพ์ใหญ่เป็น public ส่วนตัวพิมพ์เล็กใช้ได้ใน package เท่านั้น"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

field ที่ขึ้นต้นด้วยตัวพิมพ์ใหญ่เป็น exported คือ package อื่นอ่านเขียนได้ ส่วนตัวพิมพ์เล็กมองเห็นได้เฉพาะใน package ที่ประกาศ

## ผลกับการ encode

```go title="account.go"
type account struct {
	ID       int    // exported: package อื่นและ encoding/json มองเห็น
	password string // unexported: ใช้ได้แค่ใน package
}

b, _ := json.Marshal(account{ID: 1, password: "x"})
fmt.Println(string(b)) // {"ID":1}
```

`encoding/json`, `encoding/xml` และ ORM อยู่คนละ package จึงเห็นแค่ field ที่ export

## Encapsulation

ถ้าต้องป้องกันความถูกต้องของข้อมูล ให้ field เป็น unexported แล้วเปิด method แทน (`func (a *account) SetPassword(p string) error`)

:::caution[ข้อควรระวัง]
- "ทำไม JSON ออกมาเป็น `{}`" เกือบทุกครั้งเป็นเพราะ field ขึ้นต้นด้วยตัวพิมพ์เล็ก
:::
