---
title: "crypto Package"
description: "hash, HMAC และเลขสุ่มที่ปลอดภัย"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

package `crypto/...` มี primitive ที่ผ่านการตรวจสอบมาแล้ว ใช้เป็นชิ้นส่วนประกอบ อย่าคิดวิธีเข้ารหัสเอง

## ตัวอย่าง

```go title="crypto.go"
sum := sha256.Sum256([]byte("hello"))
hex.EncodeToString(sum[:]) // 2cf24dba5fb0a30e...

mac := hmac.New(sha256.New, []byte("secret")) // เซ็น webhook, token
mac.Write([]byte("msg"))
sig := mac.Sum(nil)
hmac.Equal(sig, expected) // เทียบแบบใช้เวลาคงที่

token := make([]byte, 16)
rand.Read(token) // crypto/rand
```

| ต้องการ | ใช้ |
| --- | --- |
| checksum / fingerprint | `crypto/sha256` |
| ตรวจลายเซ็นด้วย secret ที่แชร์กัน | `crypto/hmac` |
| token, ID, key | `crypto/rand` |
| เก็บรหัสผ่าน | `golang.org/x/crypto/bcrypt` หรือ `argon2` |
| เข้ารหัสข้อมูล | `crypto/aes` + `cipher.NewGCM` |
| TLS | `crypto/tls` (`net/http` ใช้อยู่) |

:::caution[ข้อควรระวัง]
- ห้าม hash รหัสผ่านด้วย SHA-256 เพราะเร็วเกินไป ให้ใช้ bcrypt หรือ argon2
- ห้ามใช้ `math/rand` กับข้อมูลลับ
- เทียบ MAC และ token ด้วย `hmac.Equal`/`subtle.ConstantTimeCompare` ไม่ใช่ `==`
:::
