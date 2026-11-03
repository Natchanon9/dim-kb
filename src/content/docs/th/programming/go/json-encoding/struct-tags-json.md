---
title: "Struct Tags สำหรับ JSON"
description: "เปลี่ยนชื่อ ข้าม และปรับรูปแบบ field ด้วย json tag"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

tag `json` กำหนดว่าแต่ละ field ที่ export จะ encode และ decode อย่างไร

## ตัวอย่าง

```go title="product.go"
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price,omitempty"` // ตัดทิ้งถ้าเป็นศูนย์
	Stock    int     `json:"stock,string"`    // encode เป็น string ของ JSON
	Internal string  `json:"-"`               // ไม่ encode เลย
}

json.Marshal(Product{ID: 1, Name: "pen", Stock: 5, Internal: "x"})
// {"id":1,"name":"pen","stock":"5"}
```

| Tag | ผล |
| --- | --- |
| `json:"name"` | ชื่อ key |
| `json:",omitempty"` | ข้ามค่าว่าง (`0`, `""`, `nil`, slice/map ว่าง) |
| `json:",omitzero"` | ข้าม zero value รวมถึง struct ศูนย์และ `time.Time` (Go 1.24+) |
| `json:",string"` | ตัวเลขหรือ bool เป็น string ของ JSON |
| `json:"-"` | ไม่สน field นี้ |

:::caution[ข้อควรระวัง]
- `omitempty` ไม่ข้าม struct ว่างหรือ `time.Time` ศูนย์ ให้ใช้ pointer หรือ `omitzero`
- ถ้าไม่ใส่ tag key จะเป็นชื่อ field ของ Go (`"ID"`) ซึ่งดูแปลกใน API
:::
