---
title: "เรียก C จาก Go"
description: "ส่ง string และ memory ข้ามขอบเขตอย่างปลอดภัย"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Go กับ C จัดการ memory คนละแบบ string และ buffer ต้อง copy ลง memory ของ C และเราต้องคืนเอง

## ส่ง string

```go title="strlen.go"
/*
#include <stdlib.h>
#include <string.h>
static size_t c_strlen(const char *s) { return strlen(s); }
*/
import "C"

func CStrlen(s string) int {
	cs := C.CString(s) // สำเนาที่ malloc ไว้ GC ของ Go ไม่ดูแล
	defer C.free(unsafe.Pointer(cs))
	return int(C.c_strlen(cs))
}
// CStrlen("hello, C") → 8
```

| การแปลง | Function |
| --- | --- |
| Go string → C | `C.CString` (ต้อง free!) |
| C string → Go | `C.GoString` |
| Go []byte → C | `C.CBytes` (ต้อง free!) |
| C buffer → Go | `C.GoBytes(ptr, n)` |
| link library | `// #cgo LDFLAGS: -lsqlite3` |

:::caution[ข้อควรระวัง]
- C เก็บ pointer ของ Go ไว้หลังการเรียกจบไม่ได้ (กฎ pointer ของ cgo) ให้ copy ข้อมูลแทน
- ทุก `C.CString`/`C.CBytes` ที่ไม่มี `C.free` คือ memory leak
:::
