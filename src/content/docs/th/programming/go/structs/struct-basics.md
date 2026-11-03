---
title: "พื้นฐาน Struct"
description: "รวม field ที่เกี่ยวข้องกันเป็น type ที่มีชื่อ"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

struct รวม field ที่มีชื่อไว้เป็น type เดียว เป็นวิธีหลักในการออกแบบข้อมูลของ Go เพราะไม่มี class

## ประกาศและใช้งาน

```go title="user.go"
type User struct {
	Name string
	Age  int
}

var u User // ทุก field เป็น zero value
u.Name = "Ann"
fmt.Println(u, u.Age) // {Ann 0} 0
```

เข้าถึง field ด้วย `.` ได้ทั้งบนค่าและ pointer (`p.Name` ใช้ได้เมื่อ `p` เป็น `*User`)

## เป็นค่า (value semantics)

กำหนดค่าหรือส่ง struct จะ copy ทุก field ถ้า function ต้องแก้ค่าหรือ struct ใหญ่ ให้ใช้ pointer (`*User`)

:::caution[ข้อควรระวัง]
- แก้สำเนาไม่ได้แก้ตัวจริง: `for _, u := range users { u.Age++ }` ไม่เปลี่ยนอะไรใน `users`
:::
