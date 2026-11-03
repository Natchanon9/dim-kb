---
title: "Code Coverage"
description: "ดูว่า test รันบรรทัดไหนบ้าง"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

coverage บอกว่า test รัน statement ไหนบ้าง ใช้หา branch ที่ยังไม่มี test ไม่ใช่ใช้เป็นคะแนนที่ต้องดันให้สูงสุด

## คำสั่ง

```sh
go test -cover ./...                          # เปอร์เซ็นต์ราย package
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out                 # เปิดดูแบบไฮไลต์
go tool cover -func=cover.out                 # ตารางราย function
go test -coverpkg=./... -coverprofile=c.out ./... # นับการเรียกข้าม package
```

## ตัวอย่าง

```go title="classify_test.go"
func TestClassify(t *testing.T) {
	for in, want := range map[int]string{-1: "negative", 0: "zero", 5: "positive"} {
		if got := Classify(in); got != want {
			t.Errorf("Classify(%d) = %q, want %q", in, got, want)
		}
	}
}
// coverage: 100.0% of statements
```

ตั้งแต่ Go 1.20 `go build -cover` ใส่ instrument ให้ทั้ง binary ได้ เพื่อวัด coverage ของ integration test

:::caution[ข้อควรระวัง]
- coverage 100% ไม่ได้แปลว่าถูก บรรทัดหนึ่งอาจถูกรันโดยไม่มีใครเช็คผลลัพธ์
:::
