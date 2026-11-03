---
title: "การ Import Package"
description: "import path, การจัดกลุ่ม, alias และ blank import"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

import path คือ module path ต่อด้วยโฟลเดอร์ ส่วนในไฟล์เรียก package ด้วยชื่อของมัน

## รูปแบบ

```go title="imports.go"
import (
	"fmt"
	str "strings" // alias

	g "github.com/me/myapp/greet"
	_ "github.com/lib/pq" // ใช้แค่ side effect (ลงทะเบียน driver)
)

fmt.Println(str.Repeat("go", 2), g.Hello("alias")) // gogo Hello, alias
```

| รูปแบบ | ใช้เมื่อ |
| --- | --- |
| `import "pkg"` | ปกติ |
| `import alias "pkg"` | ชื่อชนกัน (`crand "crypto/rand"`) หรือชื่อยาว |
| `import _ "pkg"` | ให้ `init` ของ package ทำงานอย่างเดียว |
| `import . "pkg"` | เลี่ยง เพราะไม่รู้ว่าชื่อมาจากไหน |

`goimports` (หรือ editor) เพิ่ม ลบ และจัดกลุ่ม import ให้: standard library ก่อน แล้วตามด้วย third-party

:::caution[ข้อควรระวัง]
- import ที่ไม่ได้ใช้ compile ไม่ผ่าน
:::
