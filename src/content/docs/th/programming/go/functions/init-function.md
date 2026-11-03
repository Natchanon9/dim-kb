---
title: "init Function"
description: "setup ของ package ที่รันก่อน main"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`func init()` รันอัตโนมัติครั้งเดียวต่อ package หลังกำหนดค่าตัวแปรของ package และก่อน `main` เรียกเองไม่ได้

## ตัวอย่าง

```go title="config.go"
var config map[string]string

func init() {
	config = map[string]string{"env": "dev"}
}
// config["env"] → dev
```

## ลำดับ

1. package ที่ import มา initialize ก่อน
2. จากนั้นตัวแปรระดับ package แล้วตามด้วย `init` ทุกตัวใน package (ไฟล์หนึ่งมีได้หลายตัว)
3. แล้วจึง `main`

:::caution[ข้อควรระวัง]
- ทำ `init` ให้เล็ก side effect ที่ซ่อนอยู่ (เปิด DB, อ่าน env) ทำให้ test ยาก ให้ทำใน `main` แล้วส่งต่อเข้าไป
- `init` คืน error ไม่ได้ ถ้าพังต้อง `panic` หรือ `log.Fatal`
:::
