---
title: "JSON Streaming"
description: "decode และ encode JSON ตรงจาก reader และ writer"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`json.NewDecoder(r)` อ่าน JSON จาก `io.Reader` อะไรก็ได้ ส่วน `json.NewEncoder(w)` เขียนลง `io.Writer` อะไรก็ได้ ใช้กับ HTTP body, ไฟล์ และ stream ของค่าหลายตัว

## ตัวอย่าง

```go title="stream.go"
dec := json.NewDecoder(strings.NewReader(`{"id":1} {"id":2}`))
for {
	var e struct{ ID int }
	err := dec.Decode(&e)
	if errors.Is(err, io.EOF) {
		break
	}
	if err != nil {
		return err
	}
	fmt.Println(e.ID) // 1 แล้วก็ 2
}

json.NewEncoder(os.Stdout).Encode(map[string]string{"status": "ok"})
// {"status":"ok"} พร้อมขึ้นบรรทัดใหม่
```

## ใน HTTP handler

```go
dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
dec.DisallowUnknownFields()
if err := dec.Decode(&req); err != nil { ... }
```

:::caution[ข้อควรระวัง]
- `Decode` อ่านค่าเดียว ขยะที่ตามหลังไม่ถือเป็น error เว้นแต่จะเช็ค `dec.More()` หรือ decode อีกรอบ
:::
