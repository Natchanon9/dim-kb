---
title: "Error Handling Patterns"
description: "เช็คทันที return ก่อน และเติม context ชั้นละครั้ง"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

รูปแบบของ Go คือ: เช็ค error ทันทีหลังเรียก เติม context แล้ว return ก่อน ทางหลักของโค้ดจะอยู่ชิดซ้ายเสมอ

## ตัวอย่าง

```go title="config.go"
func loadConfig(path string) (string, error) {
	data, err := readFile(path)
	if err != nil {
		return "", fmt.Errorf("load config: %w", err)
	}
	return data, nil
}
// loadConfig("") → load config: empty path
```

## หลักการ

- **เช็คทุก error** จะไม่สนก็ต้องตั้งใจ ใช้ `_` และเขียน comment บอก
- **เติม context ระหว่างทางขึ้น:** บอกว่าชั้นนี้กำลังทำอะไร ไม่ใช่บอกว่าข้างล่างพังเพราะอะไร
- **จัดการครั้งเดียว:** log หรือ return อย่างใดอย่างหนึ่ง ไม่อย่างนั้น error เดียวจะถูก log ทุกชั้น
- **ตัดสินใจที่ชั้นบนสุด:** HTTP handler หรือ `main` แปลง error เป็น status code ข้อความ หรือการ exit

:::caution[ข้อควรระวัง]
- `if err != nil { return err }` ทุกที่โดยไม่เติม context ทำให้ไล่ error อย่าง `EOF` ไม่ได้เลย
:::
