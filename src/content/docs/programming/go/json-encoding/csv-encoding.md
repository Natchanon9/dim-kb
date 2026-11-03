---
title: "CSV Encoding"
description: "Read and write CSV with encoding/csv"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`encoding/csv` handles quoting and commas inside fields, which a plain `strings.Split` gets wrong.

## Example

```go title="csv.go"
r := csv.NewReader(strings.NewReader("name,age\nann,30\n\"b, c\",40\n"))
rows, _ := r.ReadAll()
fmt.Println(len(rows), rows[2][0]) // 3 b, c

var sb strings.Builder
w := csv.NewWriter(&sb)
w.Write([]string{"x", "y,z"})
w.Flush() // x,"y,z"
```

For large files, call `r.Read()` in a loop instead of `ReadAll`. Set `r.Comma = ';'` for semicolon files, and `r.FieldsPerRecord = -1` to allow uneven rows.

:::caution[Pitfalls]
- Always `w.Flush()` and check `w.Error()`.
- Excel-exported CSV may start with a UTF-8 BOM; strip it from the first header.
:::
