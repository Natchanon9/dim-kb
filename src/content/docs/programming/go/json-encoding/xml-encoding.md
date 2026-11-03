---
title: "XML Encoding"
description: "Marshal and unmarshal XML with encoding/xml"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`encoding/xml` works like `encoding/json`, with tags for element names and attributes. You'll meet it in SOAP APIs, RSS feeds and older systems.

## Example

```go title="book.go"
type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
}

b, _ := xml.Marshal(Book{ID: "1", Title: "Go"})
// <book id="1"><title>Go</title></book>
```

| Tag | Effect |
| --- | --- |
| `xml:"name"` | Element name |
| `xml:"id,attr"` | Attribute |
| `xml:",chardata"` | Text content |
| `xml:"a>b"` | Nested path |
| `xml:",omitempty"` | Skip when empty |

`xml.NewDecoder(r).Token()` streams large documents element by element.

:::caution[Pitfalls]
- Add `xml.Header` yourself if you need the `<?xml ...?>` declaration.
:::
