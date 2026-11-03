---
title: "log/slog"
description: "structured logging ใน standard library (Go 1.21+)"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`log/slog` เขียน log แบบมีโครงสร้าง: ข้อความพร้อม attribute แบบ key-value ออกเป็น text หรือ JSON ระบบ log จึงกรองตาม field ได้

## ตัวอย่าง

```go title="logger.go"
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

logger.Info("order created", "id", 42, "total", 99.5)
// {"time":"...","level":"INFO","msg":"order created","id":42,"total":99.5}

logger.With("request_id", "r1").Warn("slow query", slog.Duration("took", 2*time.Second))
// {...,"level":"WARN","msg":"slow query","request_id":"r1","took":2000000000}

logger.Debug("hidden") // ต่ำกว่า level: ไม่ถูกเขียน
```

## นิสัยที่ดี

- สร้าง logger ตัวเดียวใน `main` แล้วส่งต่อลงไป หรือตั้งด้วย `slog.SetDefault`
- `logger.With(...)` แนบ field ของ request ไว้ครั้งเดียว
- ใช้ `InfoContext(ctx, ...)` เพื่อให้ handler ดึง trace ID จาก context ได้
- production ใช้ JSON handler ส่วนในเครื่องใช้ text handler

:::caution[ข้อควรระวัง]
- จำนวน argument แบบ key-value เป็นเลขคี่จะได้ field `!BADKEY` `go vet` จับได้
- ห้าม log รหัสผ่าน token หรือเลขบัตรเต็มๆ
:::
