---
title: "bufio Package"
description: "อ่านทีละบรรทัดแบบมี buffer และเขียนแบบมี buffer"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`bufio` ห่อ reader หรือ writer ด้วย buffer การอ่านเขียนเล็กๆ หลายครั้งจึงกลายเป็น system call ใหญ่ไม่กี่ครั้ง

## อ่านทีละบรรทัด

```go title="scan.go"
sc := bufio.NewScanner(strings.NewReader("a\nb\nc")) // หรือ *os.File
n, last := 0, ""
for sc.Scan() {
	n++
	last = sc.Text()
}
fmt.Println(n, last, sc.Err()) // 3 c <nil>
```

`sc.Split(bufio.ScanWords)` แยกทีละคำแทนทีละบรรทัด

## เขียนแบบมี buffer

```go title="write.go"
var sb strings.Builder
w := bufio.NewWriter(&sb)
w.WriteString("buffered")
sb.Len() // 0: ยังอยู่ใน buffer
w.Flush()
sb.String() // buffered
```

:::caution[ข้อควรระวัง]
- `Scanner` ไม่รับบรรทัดที่ยาวเกิน 64 KB โดยค่าเริ่มต้น ขยายด้วย `sc.Buffer(buf, max)`
- ลืม `Flush()` แล้ว output ก้อนสุดท้ายหาย
:::
