---
title: "JSON กับ Struct"
description: "marshal และ unmarshal struct ด้วย encoding/json"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`encoding/json` จับคู่ field ที่ export ของ struct กับ key ของ JSON และใช้ tag เปลี่ยนชื่อหรือข้าม field

## Encode และ decode

```go title="order.go"
type Order struct {
	ID    int      `json:"id"`
	Items []string `json:"items"`
	Note  string   `json:"note,omitempty"`
}

data, _ := json.Marshal(Order{ID: 1, Items: []string{"pen"}})
// {"id":1,"items":["pen"]}

var o Order
err := json.Unmarshal([]byte(`{"id":2,"items":["a","b"],"extra":true}`), &o)
// o.ID=2, o.Items=[a b], err=nil ("extra" ที่ไม่รู้จักถูกข้าม)
```

ตอน decode จับคู่ key แบบไม่สนตัวพิมพ์ใหญ่เล็ก field ที่ไม่มีใน JSON จะเป็น zero value

:::caution[ข้อควรระวัง]
- `Unmarshal` ต้องรับ pointer (`&o`)
- field ธรรมดาแยกไม่ออกว่า "ไม่ได้ส่งมา" หรือ "ส่งค่าศูนย์มา" ถ้าสำคัญให้ใช้ field แบบ pointer (`*int`)
- ใน API ใช้ `json.NewDecoder(r.Body)` กับ `DisallowUnknownFields()` เพื่อปฏิเสธ field ที่ไม่คาดไว้
:::
