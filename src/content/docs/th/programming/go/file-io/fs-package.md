---
title: "fs Package"
description: "interface fs.FS สำหรับ tree ของไฟล์แบบอ่านอย่างเดียว"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`io/fs` ประกาศ `fs.FS` ซึ่งเป็น interface ของ filesystem แบบอ่านอย่างเดียว โค้ดที่รับ `fs.FS` ใช้ได้กับดิสก์จริง (`os.DirFS`), ไฟล์ที่ embed (`embed.FS`), ไฟล์ zip และข้อมูลใน memory สำหรับ test

## ตัวอย่าง

```go title="fsys.go"
var fsys fs.FS = fstest.MapFS{ // อยู่ใน memory ใช้สะดวกตอน test
	"a.txt":     {Data: []byte("A")},
	"dir/b.txt": {Data: []byte("B")},
}
fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
	if !d.IsDir() {
		fmt.Println(p) // a.txt, dir/b.txt
	}
	return nil
})
data, _ := fs.ReadFile(fsys, "dir/b.txt") // B
```

| แหล่ง | ได้ `fs.FS` จาก |
| --- | --- |
| ดิสก์ | `os.DirFS("templates")` |
| ใน binary | `embed.FS` |
| ใน test | `fstest.MapFS` |

:::caution[ข้อควรระวัง]
- path ของ `fs.FS` ใช้ `/` เสมอและเป็นแบบสัมพัทธ์: ไม่มี `/` นำหน้า และไม่มี `..`
:::
