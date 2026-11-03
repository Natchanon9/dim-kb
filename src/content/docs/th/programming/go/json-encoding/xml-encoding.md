---
title: "XML Encoding"
description: "marshal และ unmarshal XML ด้วย encoding/xml"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`encoding/xml` ใช้เหมือน `encoding/json` มี tag สำหรับชื่อ element และ attribute จะเจอใน SOAP API, RSS feed และระบบเก่า

## ตัวอย่าง

```go title="book.go"
type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
}

b, _ := xml.Marshal(Book{ID: "1", Title: "Go"})
// <book id="1"><title>Go</title></book>
```

| Tag | ผล |
| --- | --- |
| `xml:"name"` | ชื่อ element |
| `xml:"id,attr"` | attribute |
| `xml:",chardata"` | ข้อความข้างใน |
| `xml:"a>b"` | path ซ้อน |
| `xml:",omitempty"` | ข้ามถ้าว่าง |

`xml.NewDecoder(r).Token()` อ่านเอกสารใหญ่ทีละ element แบบ stream ได้

:::caution[ข้อควรระวัง]
- ถ้าต้องการ `<?xml ...?>` ให้ใส่ `xml.Header` เอง
:::
