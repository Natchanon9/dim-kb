---
title: "ทำไมต้อง Go?"
description: "Go เก่งอะไร ไม่เหมาะกับอะไร และแลกอะไรไปบ้าง"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Go ยอมเสียความยืดหยุ่นของภาษา เพื่อแลกกับความเรียบง่าย build เร็ว และ deploy ง่าย ซึ่งเหมาะมากกับงาน server และ tooling

## จุดแข็ง

- **feedback เร็ว:** โปรเจกต์ใหญ่ compile ได้ในไม่กี่วินาที
- **ภาษาเล็ก:** keyword 25 ตัว และ format มีแบบเดียว (`gofmt`) โค้ดทุกที่จึงหน้าตาเหมือนกัน
- **Concurrency:** goroutine และ channel อยู่ในตัวภาษา
- **Deploy:** binary ไฟล์เดียว cross-compile ด้วย `GOOS`/`GOARCH`
- **Stdlib:** HTTP server, JSON, crypto และ testing พร้อมใช้ production

## สิ่งที่แลกกัน

| เสียไป | ได้มา |
| --- | --- |
| class และ inheritance | struct, interface, composition |
| exception | error เป็นค่า เช็คด้วย `if err != nil` |
| คุม memory เอง | garbage collection |
| macro, operator overloading | โค้ดที่ทุกคนในทีมอ่านออก |

## Go เหมาะกับงานไหน

| เหมาะ | ไม่เหมาะ |
| --- | --- |
| REST/gRPC API, microservice | GUI และ mobile app แบบ native |
| CLI และงาน infra (Docker, Kubernetes, Terraform) | ML และ data science |
| Proxy, network service | Kernel, driver, hard real-time |

:::caution[ข้อควรระวัง]
- error handling ที่ยาวเป็นความตั้งใจของภาษา การห่อด้วย helper แปลกๆ มักทำให้อ่านยากขึ้น
:::
