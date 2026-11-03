---
title: "Request Context"
description: "r.Context() สำหรับการยกเลิกและค่าที่ผูกกับ request"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

ทุก request มี context ติดมา มันถูกยกเลิกเมื่อ client ตัดการเชื่อมต่อ, stream ของ HTTP/2 ถูก reset หรือ server กำลัง shutdown

## ตัวอย่าง

```go title="context.go"
type ridKey struct{}

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ridKey{}, "req-7")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func report(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(time.Second): // งานที่ช้า
		fmt.Fprint(w, "done")
	case <-r.Context().Done():
		log.Println(r.Context().Value(ridKey{}), "stopped:", r.Context().Err())
	}
}
// client ตัดการเชื่อมต่อ → req-7 stopped: context canceled
```

ส่ง `r.Context()` ไปกับทุก query ของ DB และทุกการเรียกออกไปข้างนอก request ที่ถูกทิ้งจะได้หยุดใช้ resource

:::caution[ข้อควรระวัง]
- อย่าใช้ `r.Context()` กับงานที่ต้องทำต่อหลังตอบไปแล้ว เพราะมันถูกยกเลิกทันทีที่ handler return
:::
