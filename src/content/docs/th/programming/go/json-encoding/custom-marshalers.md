---
title: "Custom Marshalers"
description: "กำหนด JSON ของ type เองด้วย MarshalJSON และ UnmarshalJSON"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

type ที่ implement `json.Marshaler` / `json.Unmarshaler` กำหนดรูปแบบ JSON ของตัวเองได้ ใช้กับเงิน, enum, รูปแบบวันที่เฉพาะ

## ตัวอย่าง: เงินเก็บเป็นสตางค์ แต่ JSON เป็น "19.99"

```go title="money.go"
type Money int64 // สตางค์

func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d.%02d"`, m/100, m%100)), nil
}

func (m *Money) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	var whole, frac int64
	if _, err := fmt.Sscanf(s, "%d.%d", &whole, &frac); err != nil {
		return fmt.Errorf("money %q: %w", s, err)
	}
	*m = Money(whole*100 + frac)
	return nil
}
// 1999 สตางค์ ↔ {"amount":"19.99"}
```

ค่าที่เป็นแบบ string ใช้ `MarshalText`/`UnmarshalText` ง่ายกว่า และใช้ได้กับ key ของ map และ encoder ตัวอื่นด้วย

:::caution[ข้อควรระวัง]
- `UnmarshalJSON` ต้องเป็น pointer receiver ส่วน `MarshalJSON` แบบ value receiver ใช้ได้ทั้งกับค่าและ pointer
- เรียก `json.Marshal(m)` ข้างใน `m.MarshalJSON` จะวนไม่จบ ให้แปลงเป็น underlying type ก่อน
:::
