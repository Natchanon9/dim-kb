---
title: "Pointer Receivers"
description: "method ที่แก้ค่า receiver"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

method ที่ receiver เป็น `*T` ทำงานกับค่าตัวจริง สิ่งที่แก้จึงติดอยู่ และ Go ใส่ `&` หรือ `*` ให้เองตอนเรียก

## ตัวอย่าง

```go title="account.go"
type Account struct{ Balance int }

func (a *Account) Deposit(n int) { a.Balance += n }

acc := Account{}
acc.Deposit(50) // Go ส่ง &acc ให้เอง
ptr := &acc
ptr.Deposit(25)
fmt.Println(acc.Balance) // 75
```

การใส่ `&` อัตโนมัติใช้ได้กับค่าที่มี address เท่านั้น: ตัวแปร, element ของ slice, field ของ struct ใช้ไม่ได้กับ element ของ map หรือค่าที่ function คืนมา (`getAccount().Deposit(1)` จะพังถ้า `getAccount` คืนเป็นค่า)

:::caution[ข้อควรระวัง]
- ผสม value กับ pointer receiver ใน type เดียวทำให้สับสน ดูการเปรียบเทียบที่ [Value vs Pointer Receiver](/dim-kb/th/programming/go/methods-interfaces/value-vs-pointer-receivers/)
:::
