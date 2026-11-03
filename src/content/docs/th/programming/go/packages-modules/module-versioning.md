---
title: "Module Versioning"
description: "semantic version, การอัปเกรด และกฎ /v2"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

module ใช้ semantic version `vMAJOR.MINOR.PATCH` Go เลือกเวอร์ชัน **ต่ำสุด** ที่ตรงกับทุกเงื่อนไข (MVS) build จึงไม่เปลี่ยนเองเมื่อมีเวอร์ชันใหม่ออก

## เลือกเวอร์ชัน

```sh
go get github.com/foo/bar@v1.4.2  # เวอร์ชันที่ระบุ
go get github.com/foo/bar@latest  # release ล่าสุด
go get -u ./...                    # อัปเกรด dependency ทั้งตรงและอ้อม (minor/patch)
go get github.com/foo/bar@none    # ลบออก
```

## Major version

ตั้งแต่ v2 ขึ้นไป major version เป็นส่วนหนึ่งของ import path: `github.com/foo/bar/v2` v1 กับ v2 อยู่ร่วมใน build เดียวกันได้

## คำศัพท์อื่น

- **Pseudo-version** (`v0.0.0-20240101120000-abcdef123456`): commit ที่ไม่มี tag
- **`replace`** ใน `go.mod`: ชี้ dependency ไปที่ fork หรือโฟลเดอร์ในเครื่อง
- **`// indirect`**: dependency ของ dependency ไม่ได้ถูกโค้ดเราใช้ตรงๆ

:::caution[ข้อควรระวัง]
- `v0.x` ไม่สัญญาเรื่อง compatibility อัปเกรด minor ก็อาจพังได้
- อย่า commit บรรทัด `replace` ที่ชี้ไปที่ดิสก์ในเครื่องตัวเอง
:::
