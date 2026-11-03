---
title: "Use Case ของ Reflection"
description: "ที่ไหน reflection คุ้มกับต้นทุน"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

reflection เหมาะกับ infrastructure แบบทั่วไปที่ต้องรับมือกับ type ที่ไม่รู้ตอน compile

## Encoder ตัวจิ๋ว

```go title="tomap.go"
func toMap(v any) map[string]any {
	val := reflect.ValueOf(v)
	t := val.Type()
	out := map[string]any{}
	for i := range t.NumField() {
		key := t.Field(i).Tag.Get("json")
		if key == "" {
			key = t.Field(i).Name
		}
		out[key] = val.Field(i).Interface()
	}
	return out
}
// toMap(struct{ID int `json:"id"`; Name string `json:"name"`}{1, "go"}) → map[id:1 name:go]
```

## ใครใช้จริง

| ด้าน | ตัวอย่าง |
| --- | --- |
| Encoding | `encoding/json`, `encoding/xml`, YAML |
| Validation | `go-playground/validator` |
| ORM / จับคู่ SQL | GORM, sqlx |
| Testing | `reflect.DeepEqual`, `go-cmp` |
| โหลด config | library ที่แปลง env/flag → struct |

:::caution[ข้อควรระวัง]
- ถ้ากำลังเขียน reflection ใน business logic เกือบทุกครั้ง interface หรือ generics คือคำตอบที่ดีกว่า
:::
