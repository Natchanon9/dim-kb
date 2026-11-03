---
title: "Go คืออะไร"
description: "จุดเด่นของ Go และเอาไปใช้ทำอะไร"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Go เป็นภาษา compiled, statically typed และมี garbage collector จาก Google เปิดเป็น open source ตั้งแต่ปี 2009 ออกแบบมาให้โค้ด server ขนาดใหญ่ compile เร็วและอ่านง่าย

## คุณสมบัติหลัก 3 ข้อ

| คุณสมบัติ | ความหมาย |
| --- | --- |
| **Compiled** | build เป็น machine code ตอนรันไม่ต้องมี VM |
| **Statically typed** | type ผิดเจอตั้งแต่ตอน compile |
| **Garbage-collected** | คืน memory ให้อัตโนมัติ |

## อะไรทำให้ Go ต่างจากภาษาอื่น

### เรียบง่ายโดยตั้งใจ

มี keyword แค่ 25 ตัว loop มีแค่ `for` และไม่มี class หรือ inheritance ใช้ struct กับ interface แทนได้โดยไม่ต้องเขียนพิธีเยอะ

### Concurrency อยู่ในตัวภาษา

goroutine คือ function ที่รันไปพร้อมกับส่วนอื่นของโปรแกรม stack เริ่มต้นแค่ไม่กี่ KB รันพร้อมกันเป็นพันตัวก็เป็นเรื่องปกติ และส่งข้อมูลหากันผ่าน channel

### Stdlib ครบ

HTTP, JSON, crypto และ testing มากับภาษาเลย ทำ JSON API ได้โดยไม่ต้องใช้ framework:

```go title="hello.go"
mux := http.NewServeMux()
mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello " + r.PathValue("name"),
	})
})
http.ListenAndServe(":8080", mux)
```

`GET /hello/gopher` ได้ `{"message":"hello gopher"}` ส่วน `POST` ได้ `405` ทันทีเพราะ method อยู่ใน pattern แล้ว (Go 1.22+)

### Binary ไฟล์เดียว

`go build` ได้ executable ไฟล์เดียวที่รันได้เลย cross-compile ก็แค่ตั้ง env ตัวเดียว: `GOOS=linux go build`

## Explicit ดีกว่า clever

format โค้ดมีแบบเดียว (`gofmt`) error เป็นค่าธรรมดาที่เช็คด้วย `if err != nil` ไม่มี exception, operator overloading หรือ macro โค้ดทุกทีมจึงหน้าตาเหมือนกัน แลกกับการพิมพ์ยาวขึ้นนิดหน่อย

## Go เหมาะกับงานไหน

| เหมาะ | ไม่เหมาะ |
| --- | --- |
| REST API, microservice | GUI และ mobile app แบบ native |
| CLI และงาน infra (Docker, Kubernetes, Terraform) | ML และ data science |
| Proxy และ network service | Kernel, driver, hard real-time |

:::caution[ข้อควรระวัง]
- อย่าเขียน Go แบบ Java: ไม่ต้องมี class hierarchy หรือ getter/setter ทุก field
- goroutine เบาแต่ไม่ฟรี ตัวที่ไม่จบ = memory leak
- binary ใหญ่กว่า C/C++ เพราะรวม runtime และ GC มาด้วย
:::
