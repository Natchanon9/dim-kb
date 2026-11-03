---
title: "Cloud Run + Cloud SQL"
description: "deploy container ของ Go ขึ้น Cloud Run และต่อ Cloud SQL Postgres"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Cloud Run รัน container ตามปริมาณ request ลดเหลือศูนย์ได้ และมี HTTPS ให้ แอปแค่ต้อง listen ที่ `$PORT` และจัดการ SIGTERM

## สิ่งที่แอปต้องมี

```go title="main.go"
port := os.Getenv("PORT") // Cloud Run ตั้งให้ ส่วนใหญ่เป็น 8080
if port == "" {
	port = "8080"
}
srv := &http.Server{Addr: ":" + port, Handler: router}
// และ graceful shutdown เมื่อได้ SIGTERM (ดูหมวด Networking และ HTTP)
```

## Deploy

```sh
gcloud run deploy myapp \
  --source . \
  --region asia-southeast1 \
  --allow-unauthenticated \
  --add-cloudsql-instances PROJECT:asia-southeast1:mydb \
  --set-env-vars DB_USER=app,DB_NAME=app \
  --set-secrets DB_PASS=db-pass:latest
```

`--source .` build ด้วย Dockerfile (หรือ buildpacks) ผ่าน Cloud Build

## ต่อ Cloud SQL

เมื่อใช้ `--add-cloudsql-instances` Cloud Run จะเปิด Unix socket ให้:

```text
host=/cloudsql/PROJECT:REGION:INSTANCE user=app password=... dbname=app
```

หรือใช้ Cloud SQL Go connector (`cloud.google.com/go/cloudsqlconn`) พร้อมยืนยันตัวตนด้วย IAM

:::caution[ข้อควรระวัง]
- จำนวน instance × ขนาด pool อาจเกิน `max_connections` ของ Postgres ให้ตั้ง `SetMaxOpenConns` น้อยๆ หรือจำกัดจำนวน instance สูงสุด
- secret ต้องอยู่ใน Secret Manager (`--set-secrets`) ไม่ใช่ env ธรรมดาหรือใน image
:::
