---
title: "พื้นฐาน JSON"
description: "Marshal และ Unmarshal ด้วย encoding/json"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`json.Marshal` แปลงค่า Go เป็น byte ของ JSON ส่วน `json.Unmarshal` เติมค่า Go จาก JSON ใช้ได้กับ struct, map, slice และ type พื้นฐาน

## ตัวอย่าง

```go title="json.go"
b, _ := json.Marshal(map[string]any{"name": "go", "tags": []string{"a"}, "ok": true})
// {"name":"go","ok":true,"tags":["a"]}  (key ของ map เรียงให้)

var v map[string]any
json.Unmarshal([]byte(`{"n":1}`), &v)
fmt.Printf("%T\n", v["n"]) // float64

pretty, _ := json.MarshalIndent(map[string]int{"a": 1}, "", "  ")
```

## การจับคู่ type

| JSON | Go (ลง `any`) | Go (มี type) |
| --- | --- | --- |
| object | `map[string]any` | struct, `map[string]T` |
| array | `[]any` | slice, array |
| number | `float64` | int/float ทุกแบบ |
| string | `string` | `string`, `time.Time`, `[]byte` (base64) |
| `true`/`false` | `bool` | `bool` |
| `null` | `nil` | pointer, slice, map → nil |

decode ลง struct ดีกว่าลง `map[string]any`

:::caution[ข้อควรระวัง]
- ID ที่เป็นตัวเลขใหญ่จะเสียความแม่นยำเมื่อเป็น `float64` ให้ decode ลง field `int64` หรือใช้ `dec.UseNumber()`
:::
