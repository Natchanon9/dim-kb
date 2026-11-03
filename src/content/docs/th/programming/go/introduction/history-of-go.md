---
title: "ประวัติของ Go"
description: "Go มาจากไหน และเวอร์ชันที่เปลี่ยนวิธีเขียน Go"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Go เริ่มที่ Google ในปี 2007 ออกแบบโดย Robert Griesemer, Rob Pike และ Ken Thompson และเปิดเป็น open source เดือนพฤศจิกายน 2009

## ทำไม Google ถึงสร้าง Go

service ที่เขียนด้วย C++ ขนาดใหญ่ build นานมาก และ codebase ใหญ่ๆ ก็อ่านยาก เป้าหมายคือภาษาที่ build เร็ว syntax เรียบง่าย และรองรับเครื่อง multicore กับงาน server บนเครือข่ายได้ดี

## เวอร์ชันที่ควรรู้

| เวอร์ชัน | ปี | เปลี่ยนอะไร |
| --- | --- | --- |
| 1.0 | 2012 | stable ตัวแรก และ **Go 1 compatibility promise** |
| 1.5 | 2015 | compiler และ runtime เขียนใหม่ด้วย Go; GC แบบ concurrent latency ต่ำ |
| 1.11 | 2018 | **Go modules** (จัดการ dependency โดยไม่ต้องใช้ GOPATH) |
| 1.13 | 2019 | error wrapping: `%w`, `errors.Is`, `errors.As` |
| 1.18 | 2022 | **Generics** |
| 1.21 | 2023 | builtin `min`/`max`, package `slices` และ `maps`, `log/slog` |
| 1.22 | 2024 | loop variable แยกต่อรอบ, `for i := range n`, ใส่ method + path pattern ใน `ServeMux` ได้ |
| 1.23 | 2024 | range over function (iterator) |

## Compatibility promise

โค้ดที่เขียนสำหรับ Go 1.0 ยัง compile ได้กับ Go 1.x ปัจจุบัน tutorial เก่าส่วนใหญ่จึงยังใช้ได้ และฟีเจอร์ใหม่ของภาษาก็มาช้า

:::caution[ข้อควรระวัง]
- tutorial เก่าอาจใช้ GOPATH mode, `ioutil` หรือ copy loop variable เอง ซึ่งล้าสมัยแล้วตั้งแต่มี modules (1.11), Go 1.16 และ Go 1.22 ตามลำดับ
:::
