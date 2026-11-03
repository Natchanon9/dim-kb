---
title: "json.RawMessage"
description: "เลื่อนการ decode บางส่วนของ JSON ไว้ทีหลัง"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`json.RawMessage` เก็บชิ้นส่วนของ JSON เป็น byte ดิบ decode ซองนอกก่อน แล้วค่อย decode เนื้อในเมื่อรู้ type แล้ว

## ตัวอย่าง

```go title="event.go"
type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

var e Event
json.Unmarshal([]byte(`{"type":"user","data":{"name":"ann"}}`), &e)
switch e.Type {
case "user":
	var u struct {
		Name string `json:"name"`
	}
	json.Unmarshal(e.Data, &u) // ann
}
```

ใช้บ่อยกับ payload ของ webhook และ message queue ที่มี field `type`, เก็บ column JSON โดยไม่ parse และส่ง JSON ผ่านไปโดยไม่แก้

:::caution[ข้อควรระวัง]
- field ที่ไม่มีทำให้ `RawMessage` เป็น nil ส่วน `null` ได้ byte คำว่า `null` ต้องเช็คทั้งสองกรณีก่อน decode
:::
