---
title: "Fuzzing"
description: "ให้ go test สร้าง input เองเพื่อหาจุดพัง (Go 1.18+)"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

fuzz test `FuzzXxx(f *testing.F)` เริ่มจาก seed input แล้วกลายพันธุ์มันไปเรื่อยๆ เพื่อหา input ที่ทำให้คุณสมบัติที่เรา assert ไว้พัง

## ตัวอย่าง

```go title="reverse_test.go"
func FuzzReverse(f *testing.F) {
	f.Add("hello") // seed corpus
	f.Add("สวัสดี")
	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip()
		}
		if got := Reverse(Reverse(s)); got != s {
			t.Errorf("Reverse twice = %q, want %q", got, s)
		}
	})
}
```

```sh
go test ./...                          # รันแค่ seed
go test -fuzz=FuzzReverse -fuzztime=30s # fuzz จริง
```

input ที่ทำให้พังจะถูกเก็บไว้ที่ `testdata/fuzz/FuzzReverse/` และถูกรันซ้ำทุกครั้งที่ `go test`

## เป้าหมายที่เหมาะ

parser, decoder, validator และอะไรก็ตามที่รับ input ที่ไม่น่าไว้ใจ ให้ assert คุณสมบัติ: แปลงไปกลับได้ค่าเดิม, ไม่ panic, invariant ที่ต้องจริงเสมอ

:::caution[ข้อควรระวัง]
- fuzz ได้ทีละ target ต่อคำสั่ง ใน CI ส่วนใหญ่แค่รัน seed
:::
