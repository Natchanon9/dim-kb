---
title: "โปรแกรม Go แรก"
description: "Hello World: package main, func main, go run และ go build"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

ทุกโปรแกรม Go เริ่มที่ `package main` ใน function `main` สร้าง module ก่อนเพื่อให้เครื่องมือรู้ว่า root ของโปรเจกต์อยู่ไหน

## สร้างและรัน

```sh
mkdir hello && cd hello
go mod init example.com/hello
```

```go title="main.go"
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

```sh
go run .        # compile + รัน ไม่เก็บไฟล์
go build        # ได้ binary ชื่อ ./hello
```

## แต่ละบรรทัดทำอะไร

- **`package main`:** บอกว่าเป็นโปรแกรมที่รันได้ ไม่ใช่ library
- **`import "fmt"`:** ดึง package สำหรับ format มาใช้ import แล้วไม่ใช้จะ compile ไม่ผ่าน
- **`func main()`:** จุดเริ่มโปรแกรม ไม่รับ argument ไม่คืนค่า ใช้ `os.Args` และ `os.Exit` แทน

:::caution[ข้อควรระวัง]
- ตัวแปรและ import ที่ไม่ได้ใช้ compile ไม่ผ่าน ลบทิ้งหรือใช้ `_`
- `{` ต้องอยู่บรรทัดเดียวกับ `func main()` เพราะ Go ใส่ semicolon ท้ายบรรทัดให้อัตโนมัติ
:::
