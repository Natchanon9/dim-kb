---
title: "Project Layout"
description: "cmd/, internal/ และการทำโครงให้แบน"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Go ไม่มีโครงโปรเจกต์บังคับ เริ่มแบบแบนๆ แล้วค่อยเพิ่มโฟลเดอร์เมื่อโค้ดต้องการ รูปแบบที่เจอบ่อยสำหรับ service:

```text
myapp/
├─ go.mod
├─ cmd/
│  ├─ api/main.go        หนึ่งโฟลเดอร์ต่อหนึ่ง binary
│  └─ worker/main.go
├─ internal/             ใช้ได้แค่ใน module นี้
│  ├─ order/             หนึ่ง feature: handler, service, repository
│  ├─ user/
│  └─ platform/          setup ของ db, config, logging
├─ migrations/           ไฟล์ SQL
└─ Dockerfile
```

## แนวทาง

- **`main` ต้องบาง:** อ่าน config ต่อ dependency แล้วเริ่ม server
- **จัดตาม feature** (`order`, `user`) ไม่ใช่ตามชั้น (`controllers`, `models`)
- **`internal/`** สำหรับทุกอย่างที่ไม่ให้ module อื่น import
- **`pkg/`** ใช้เฉพาะเมื่อตั้งใจเผยแพร่ package ให้คนอื่นใช้ หลายโปรเจกต์ไม่ต้องมี
- **โปรเจกต์เล็ก?** `main.go` กับไฟล์ไม่กี่ไฟล์ที่ root ก็พอ

:::caution[ข้อควรระวัง]
- "golang-standards/project-layout" เป็น repo ของชุมชน ไม่ใช่มาตรฐานทางการ อย่า copy ทุกโฟลเดอร์มา
:::
