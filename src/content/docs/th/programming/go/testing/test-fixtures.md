---
title: "Test Fixtures"
description: "โฟลเดอร์ testdata และ golden file"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

ใส่ไฟล์ input และ output ที่คาดไว้ในโฟลเดอร์ `testdata/` เครื่องมือ go จะไม่มองว่าเป็น package และ test รันโดยใช้โฟลเดอร์ของ package เป็น working directory

## Golden file

```go title="golden_test.go"
func TestGreetGolden(t *testing.T) {
	want, err := os.ReadFile(filepath.Join("testdata", "greeting.golden"))
	if err != nil {
		t.Fatal(err)
	}
	if got := Greet("Ann") + "\n"; got != string(want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
```

ที่นิยมเพิ่มคือ flag `-update` สำหรับเขียน golden file ใหม่เมื่อ output เปลี่ยนโดยตั้งใจ:

```go
var update = flag.Bool("update", false, "rewrite golden files")
// if *update { os.WriteFile(goldenPath, []byte(got), 0o644) }
```

เหมาะกับ template ที่ render แล้ว, โค้ดที่ generate และ JSON response ก้อนใหญ่

:::caution[ข้อควรระวัง]
- ตรวจ diff ของ golden file ให้ดี `-update` ยอมรับทุกอย่างที่โค้ดสร้างออกมาตอนนี้
:::
