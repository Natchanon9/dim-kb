---
title: "Testing Best Practices"
description: "นิสัยที่ทำให้ test ของ Go เร็วและเชื่อถือได้"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

test ที่ดีเร็ว เป็นอิสระต่อกัน และ test พฤติกรรมผ่าน interface เล็กๆ

## ทำตัวปลอมผ่าน interface

```go title="welcome_test.go"
type Mailer interface{ Send(to, body string) error }

type fakeMailer struct{ sent []string }

func (f *fakeMailer) Send(to, _ string) error {
	f.sent = append(f.sent, to)
	return nil
}

func TestWelcome(t *testing.T) {
	m := &fakeMailer{}
	if err := Welcome(m, "a@b.com"); err != nil {
		t.Fatal(err)
	}
	if len(m.sent) != 1 {
		t.Errorf("sent = %v", m.sent)
	}
}
```

## เช็คลิสต์

- **table-driven + subtest** กับอะไรที่มีหลาย case
- **เขียนตัวปลอมเอง** สำหรับ interface เล็ก ส่วน interface ใหญ่ใช้ตัว generate mock (`mockgen`)
- **ไม่ sleep** ให้รอที่ channel หรือ poll พร้อม deadline
- **ใช้ของจริงใน integration test** (testcontainers, database สำหรับ test) โดยแยกด้วย build tag หรือเช็ค `-short`
- **test ผ่าน public API** ของ package (`package foo_test`) ถ้าทำได้
- **รัน `-race`** ใน CI

:::caution[ข้อควรระวัง]
- test ที่ขึ้นกับลำดับของ map, เวลา หรือลำดับของ test จะ flaky ให้แก้ อย่าแค่ retry
:::
