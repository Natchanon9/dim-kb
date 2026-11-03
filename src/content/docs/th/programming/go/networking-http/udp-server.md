---
title: "UDP Server"
description: "datagram ด้วย net.ListenPacket"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

UDP ส่ง datagram ที่เป็นอิสระต่อกัน ไม่มี connection ไม่รับประกันลำดับหรือการส่งถึง เหมาะกับ DNS, metrics (StatsD), เกม และการค้นหาเครื่องในเครือข่าย

## ตัวอย่าง

```go title="udp.go"
pc, err := net.ListenPacket("udp", ":9001")
if err != nil {
	log.Fatal(err)
}
buf := make([]byte, 1500)
for {
	n, addr, err := pc.ReadFrom(buf) // datagram หนึ่งก้อนพร้อมผู้ส่ง
	if err != nil {
		continue
	}
	pc.WriteTo(bytes.ToUpper(buf[:n]), addr)
}
// client ส่ง "ping" → ได้ "PING"
```

| | TCP | UDP |
| --- | --- | --- |
| Connection | มี | ไม่มี |
| ลำดับ / การส่งถึง | รับประกัน | ไม่รับประกัน |
| หน่วย | stream ของ byte | datagram |

:::caution[ข้อควรระวัง]
- datagram ที่ใหญ่กว่า buffer จะถูกตัด ให้ข้อความเล็กกว่าประมาณ 1400 byte เพื่อไม่ให้ถูกแบ่ง fragment
:::
