---
title: "Go คืออะไร"
description: "จุดเด่นของ Go และเอาไปใช้ทำอะไร"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
sources:
  - https://algomaster.io/learn/go/what-is-go
  - https://algomaster.io/learn/go/why-go
  - https://mikelopster.dev/docs/c/goapi-essential/chapter-1/intro
---

## TL;DR

- Go = ภาษา **compiled + statically typed + garbage-collected** จาก Google (Griesemer, Pike, Thompson เริ่ม 2007, open source 2009)
- ออกแบบให้ **เรียบง่าย**: keyword แค่ 25 ตัว, loop มีแค่ `for`, ไม่มี class / inheritance ใช้ struct + interface แทน
- **Concurrency ในตัวภาษา**: goroutine เบามาก (stack เริ่มแค่ไม่กี่ KB) คุยกันผ่าน channel
- Build ออกมาเป็น **single binary** ไม่ต้องลง runtime/VM — deploy ง่าย, cross-compile ได้เลย
- Stdlib แน่น: HTTP server, JSON, crypto, testing พร้อมใช้ production

## แนวคิด

Go เกิดจากความหงุดหงิดของทีม Google กับ C++ ที่ build ทีละหลายสิบนาที และโค้ดสาย Java ที่ inheritance ซ้อนกันจนอ่านยาก เป้าหมายเลยคือภาษาที่ compile เร็ว อ่านง่าย และเขียนโปรแกรม server ที่รับ request พร้อมกันเยอะๆ ได้โดยไม่ต้องปวดหัว

ปรัชญาหลักคือ **explicit ดีกว่า clever**: format โค้ดมีแบบเดียว (`gofmt`), error เป็นค่าธรรมดาที่ต้อง `if err != nil` เอง ไม่มี exception, ไม่มี operator overloading หรือ macro ผลคือโค้ดของทุกคนหน้าตาคล้ายกัน เปิด repo คนอื่นแล้วอ่านรู้เรื่องเร็ว แลกกับการที่ต้องพิมพ์ยาวขึ้นนิดหน่อย

ข้อแลกเปลี่ยนที่ควรรู้: generics เพิ่งมาใน 1.18, ecosystem เล็กกว่า Java/Python, GC ทำให้คุม memory ได้น้อยกว่า Rust/C++ และไม่เหมาะกับ GUI, mobile native, ML/data science

## Code

API ตัวเล็กที่สุดด้วย stdlib ล้วน ไม่ต้องลง framework (Go 1.22+ ใส่ method และ path param ใน pattern ได้)

```go title="hello.go"
mux := http.NewServeMux()
mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello " + r.PathValue("name"),
	})
})
http.ListenAndServe(":8080", mux)
// GET  /hello/gopher -> {"message":"hello gopher"}
// POST /hello/gopher -> 405 Method Not Allowed (ได้ฟรีจาก pattern)
```

ตัวเต็ม + test: `examples/go/intro/what-is-go/`

## Gotchas

:::caution
- "เรียบง่าย" ≠ "ทำได้น้อย" แต่หมายถึงมีทางทำแค่ไม่กี่ทาง — อย่าพยายามเขียน Go ให้เหมือน Java (class hierarchy, getter/setter ทุก field)
- Goroutine เบาก็จริง แต่ไม่ได้ฟรี: ปล่อยค้างไว้ = goroutine leak กิน memory เรื่อยๆ (ดูหมวด Concurrency)
- Single binary ใหญ่กว่า C/C++ เพราะรวม runtime + GC มาด้วย
:::

## ใช้จริงเมื่อไหร่

เลือก Go เมื่อทำ REST API / microservice / CLI / งาน infra (Docker, Kubernetes, Terraform เขียนด้วย Go ทั้งหมด) — โดยเฉพาะงานที่ต้องรับ concurrent request เยอะและอยาก deploy เป็น container เล็กๆ

## อ้างอิง

- <https://algomaster.io/learn/go/what-is-go>
- <https://algomaster.io/learn/go/why-go>
- <https://mikelopster.dev/docs/c/goapi-essential/chapter-1/intro>
