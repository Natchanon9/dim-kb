---
title: "Graceful Shutdown"
description: "ทำ request ที่ค้างอยู่ให้เสร็จก่อนปิดโปรแกรม"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

เมื่อได้ SIGTERM (Kubernetes, Cloud Run, Docker stop) ให้หยุดรับ connection ใหม่ ปล่อยให้ request ปัจจุบันทำจนเสร็จ แล้วค่อยปิด `http.Server.Shutdown` ทำสิ่งนี้ให้

## ตัวอย่าง

```go title="main.go"
srv := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second}

go func() {
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}()

ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()
<-ctx.Done() // รอสัญญาณ

shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
if err := srv.Shutdown(shutdownCtx); err != nil {
	log.Println("forced shutdown:", err)
}
// แล้วปิด DB pool และ flush log
```

`ListenAndServe` คืน `http.ErrServerClosed` ทันทีหลัง `Shutdown` ซึ่งเป็นเรื่องปกติ ไม่ใช่ความล้มเหลว

:::caution[ข้อควรระวัง]
- ตั้ง timeout ของ shutdown ให้น้อยกว่า grace period ของ platform (Kubernetes ค่าเริ่มต้น 30 วินาที)
- `Shutdown` ไม่รอ connection ที่ถูก hijack หรือ WebSocket ต้องติดตามพวกนั้นเอง
:::
