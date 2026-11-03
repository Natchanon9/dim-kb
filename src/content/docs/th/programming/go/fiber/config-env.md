---
title: "Config และ ENV"
description: "โหลด config จาก environment variable ตอนเริ่มโปรแกรม"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

อ่าน config ครั้งเดียวตอนเริ่มโปรแกรมลงใน struct ตรวจความถูกต้อง แล้วส่งต่อลงไป environment variable เป็นมาตรฐานของ container (12-factor)

## ตัวอย่าง

```go title="config.go"
type Config struct {
	Port        int
	DatabaseURL string
	Timeout     time.Duration
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func LoadConfig() (Config, error) {
	port, err := strconv.Atoi(getenv("PORT", "8080"))
	if err != nil {
		return Config{}, err
	}
	timeout, err := time.ParseDuration(getenv("TIMEOUT", "5s"))
	if err != nil {
		return Config{}, err
	}
	return Config{Port: port, DatabaseURL: os.Getenv("DATABASE_URL"), Timeout: timeout}, nil
}
```

test ด้วย `t.Setenv("PORT", "9000")`

## ไฟล์ .env ในเครื่อง

`github.com/joho/godotenv` โหลดไฟล์ `.env` เข้า environment ตอนพัฒนาในเครื่อง: `godotenv.Load()` บรรทัดแรกของ `main` ส่วน library อย่าง `caarlos0/env` หรือ Viper จับคู่ env กับ tag ของ struct ให้

:::caution[ข้อควรระวัง]
- commit `.env.example` แต่ห้าม commit `.env` ที่มี secret จริง
- config ที่จำเป็นแต่ไม่มี ให้พังตั้งแต่ตอนเริ่มโปรแกรม ไม่ใช่ตอน request แรก
:::
