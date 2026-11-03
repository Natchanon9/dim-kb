---
title: "JSON APIs"
description: "decode request ตรวจข้อมูล และเขียน JSON response"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

endpoint แบบ JSON ทำ 3 อย่าง: decode body, ตรวจข้อมูล และเขียน JSON response พร้อม status ที่ถูกต้อง

## ตัวอย่าง

```go title="users.go"
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}
	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name required"})
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"id": 1, "name": req.Name})
}
// {"name":"ann"} → 201 {"id":1,"name":"ann"}
// {}             → 400 {"error":"name required"}
```

## นิสัยที่ดี

- แยก struct ของ request/response ออกจาก model ของ database
- จำกัดขนาด body ด้วย `http.MaxBytesReader`
- ทุก endpoint ใช้รูปแบบ error เดียวกัน

:::caution[ข้อควรระวัง]
- อย่าคืน error ภายใน (ข้อความ SQL) ให้ client ตรงๆ ให้ log ไว้แล้วส่งข้อความทั่วไปแทน
:::
