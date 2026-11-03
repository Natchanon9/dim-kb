---
title: "sync.Pool"
description: "ใช้ object ชั่วคราวซ้ำเพื่อลดการจอง memory"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`sync.Pool` เก็บ object สำรองไว้ใช้ซ้ำ ลดการจอง memory และงานของ GC ในส่วนที่ทำงานหนักอย่าง buffer สำหรับ encode

## ตัวอย่าง

```go title="pool.go"
var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func render(name string) string {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset() // object จาก pool อาจมีข้อมูลเก่าค้าง
	buf.WriteString("hello, ")
	buf.WriteString(name)
	return buf.String() // hello, go
}
```

## กติกา

- pool ทิ้ง object ได้ทุกครั้งที่ GC ทำงาน อย่าใช้เป็น cache
- reset object ก่อนใช้เสมอ
- วัดผลก่อน: คุ้มเฉพาะกับ object ที่จองบ่อยและอายุสั้น

:::caution[ข้อควรระวัง]
- อย่าคืน buffer ที่ใหญ่มากกลับเข้า pool มันจะถือ memory ไว้ เช็ค `buf.Cap()` แล้วทิ้งตัวที่ใหญ่เกิน
:::
