---
title: "Publish Module"
description: "tag เวอร์ชันแล้วให้ proxy กับ pkg.go.dev ดึงไป"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

การ publish Go module คือการ push git tag ไม่มีขั้นตอน upload ขึ้น registry

## ขั้นตอน

```sh
go mod tidy && go test ./...
git tag v1.2.0
git push origin v1.2.0
GOPROXY=proxy.golang.org go list -m github.com/me/lib@v1.2.0  # สั่งให้ proxy ดึงไป
```

module path ใน `go.mod` ต้องตรงกับ URL ของ repository และ doc จะขึ้นบน pkg.go.dev เมื่อ proxy เห็นเวอร์ชันนั้นแล้ว

## ดูแลให้ดี

- ทำตาม semver: breaking change → major version ใหม่และ path `/v2`
- ออก release ผิดไป? เพิ่ม `retract v1.2.0` ใน `go.mod` แล้ว publish เวอร์ชันใหม่
- เขียน doc comment ให้ทุกชื่อที่ export เพราะจะกลายเป็นหน้าบน pkg.go.dev

:::caution[ข้อควรระวัง]
- tag แทบจะถาวร เพราะ proxy cache ไว้ อย่าย้ายหรือ push tag เดิมซ้ำเด็ดขาด
:::
