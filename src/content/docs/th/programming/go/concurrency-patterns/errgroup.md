---
title: "errgroup"
description: "รัน goroutine เก็บ error แรก และยกเลิกตัวที่เหลือ"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`golang.org/x/sync/errgroup` คือ WaitGroup ที่คืน error ได้ ถ้าใช้ `WithContext` ความล้มเหลวครั้งแรกจะยกเลิก context ที่แชร์กัน goroutine อื่นจึงหยุดได้

## ตัวอย่าง

```go title="errgroup.go"
g, ctx := errgroup.WithContext(context.Background())
g.SetLimit(2) // พร้อมกันได้สูงสุด 2
for _, u := range []string{"a", "bad", "c"} {
	g.Go(func() error { return fetch(ctx, u) })
}
if err := g.Wait(); err != nil {
	fmt.Println(err) // fetch bad: failed
}
```

`fetch` ควรดู `ctx.Done()` จะได้หยุดเมื่อการเรียกตัวอื่นล้มเหลวแล้ว

## ใช้เมื่อไหร่

เรียกหลายงานพร้อมกันโดยที่พังตัวไหนก็ถือว่าพังทั้งหมด: โหลดข้อมูลจากหลาย service มาทำ response เดียว, upload ไฟล์เป็นหลายส่วน

:::caution[ข้อควรระวัง]
- `Wait` คืนแค่ error **แรก** ถ้าต้องการทั้งหมดให้รวมเองด้วย `errors.Join`
- ใน `g.Go` ให้ใช้ `ctx` ที่ได้จาก `WithContext` ไม่ใช่ตัวแม่
:::
