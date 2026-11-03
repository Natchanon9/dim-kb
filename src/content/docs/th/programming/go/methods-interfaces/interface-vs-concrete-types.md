---
title: "Interface vs Concrete Type"
description: "เมื่อไหร่ควรคืน struct และกับดัก typed nil"
sidebar:
  order: 12
status: done
lastReviewed: 2026-11-03
---

คืน concrete type เพื่อให้ผู้เรียกได้ทุก method และ field และรับ interface เพื่อให้ผู้เรียกส่งอะไรก็ได้ที่เข้ากัน จุดเดียวที่พังคือคืน pointer ที่เป็น nil ออกไปเป็น interface

## กับดัก typed nil

```go title="typednil.go"
type MyErr struct{}

func (*MyErr) Error() string { return "boom" }

func mayFail(fail bool) error {
	var p *MyErr // pointer ที่เป็น nil
	if fail {
		p = &MyErr{}
	}
	return p // BUG: interface ไม่เป็น nil แต่ข้างในถือ *MyErr ที่เป็น nil
}

err := mayFail(false)
fmt.Println(err == nil) // false
```

interface ถือ `(type=*MyErr, value=nil)` และ interface จะเป็น nil ก็ต่อเมื่อ **ทั้งสอง** ส่วนเป็น nil วิธีแก้คือ `return nil` ตรงๆ ในทางที่สำเร็จ

## เลือกแบบไหน

| คืน concrete type | คืน interface |
| --- | --- |
| constructor (`NewServer`) | `error` |
| ผู้เรียกต้องใช้ field/method | มีหลาย implementation ที่เลือกตอน runtime |

:::caution[ข้อควรระวัง]
- อย่าคืน pointer ที่เป็น nil ผ่านผลลัพธ์ที่เป็น `error` เด็ดขาด
:::
