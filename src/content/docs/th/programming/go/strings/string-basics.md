---
title: "พื้นฐาน String"
description: "byte แบบ UTF-8 ที่แก้ไม่ได้, len, index, raw string และการต่อ string"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

string ของ Go คือลำดับ byte ที่แก้ไขไม่ได้ ปกติเป็นข้อความ UTF-8 operation ที่ "แก้" string จริงๆ คือสร้าง string ใหม่

## พื้นฐาน

```go title="string.go"
s := "go"
t := s + "lang"    // สร้าง string ใหม่
raw := `C:\path\n` // raw string: ไม่ตีความ escape และขึ้นหลายบรรทัดได้
fmt.Println(t, len("héllo"), s[0], raw)
// golang 6 103 C:\path\n
```

- `len(s)` นับ **byte** ไม่ใช่ตัวอักษร: `é` ใช้ 2 byte
- `s[i]` เป็น **byte** (`uint8`) ไม่ใช่ตัวอักษร
- `s[i:j]` ตัดตาม byte และแชร์ memory กับ `s`
- เทียบ string ด้วย `==` และ `<` ได้ โดยเทียบทีละ byte

:::caution[ข้อควรระวัง]
- `s[0] = 'G'` compile ไม่ผ่าน ต้องแปลงเป็น `[]byte` หรือ `[]rune` แก้แล้วแปลงกลับ
- ตัดด้วย index ของ byte อาจตัดกลางตัวอักษรที่มีหลาย byte (ภาษาไทยใช้ 3 byte ต่อตัว)
:::
