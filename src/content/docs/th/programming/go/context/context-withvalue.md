---
title: "context.WithValue"
description: "แนบค่าที่ผูกกับ request เช่น request ID"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`WithValue` แนบคู่ key-value ไปกับ context ใช้เฉพาะข้อมูลที่ผูกกับ request และต้องข้ามขอบเขตของ API: request ID, trace span, ผู้ใช้ที่ login แล้ว

## ตัวอย่าง

```go title="value.go"
type requestIDKey struct{} // type ของ key แบบ unexported: ไม่ชนกับใคร

func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, id)
}

func requestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey{}).(string)
	return id, ok
}

ctx := withRequestID(context.Background(), "req-42")
requestID(ctx)                  // req-42 true
requestID(context.Background()) // "" false
```

ห่อการเข้าถึงด้วย helper function ที่มี type ผู้เรียกจะได้ไม่ต้องแตะ key หรือ type assertion เลย

:::caution[ข้อควรระวัง]
- อย่าส่ง parameter ของ function ผ่าน context (DB handle, config, argument ที่ไม่บังคับ) ให้ส่งตรงๆ
- อย่าใช้ `string` ธรรมดาเป็น key เพราะ package อื่นอาจใช้ชื่อชนกัน
:::
