---
title: "Struct Literals"
description: "สร้าง struct ด้วย literal แบบระบุชื่อ field, ตามลำดับ และแบบ pointer"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

struct literal สร้างค่าโดยกำหนด field บางตัวหรือทั้งหมด field ที่ไม่ใส่จะได้ zero value

## 3 รูปแบบ

```go title="literals.go"
a := User{Name: "Bob", Age: 20} // ระบุชื่อ field: แนะนำ
b := User{"Cat", 30}            // ตามลำดับ: ต้องครบทุก field
p := &User{Name: "Dan"}         // pointer ไปยัง struct ใหม่
p.Age = 40                      // เหมือน (*p).Age
fmt.Println(a, b, *p) // {Bob 20} {Cat 30} {Dan 40}
```

ใช้แบบระบุชื่อ field เถอะ เพิ่มหรือสลับ field ทีหลังก็ไม่พัง `go vet` จะเตือนถ้าใช้แบบตามลำดับกับ type ของ package อื่น

:::caution[ข้อควรระวัง]
- ถ้า `}` ปิดอยู่บรรทัดของตัวเอง ต้องมี comma ท้าย field สุดท้าย
:::
