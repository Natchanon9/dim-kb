---
title: "Docker"
description: "image เล็กและปลอดภัยด้วย multi-stage build"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Go compile เป็น static binary ไฟล์เดียว image ตอนรันจึงแทบว่างเปล่าได้ Dockerfile แบบ multi-stage จะ build ใน image ของ Go แล้ว copy แค่ binary ออกมา

## Dockerfile

```dockerfile title="Dockerfile"
FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download                 # ใช้ cache ถ้า dependency ไม่เปลี่ยน
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /app ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
```

```sh
docker build -t myapp .
docker run -p 8080:8080 -e DATABASE_URL=... myapp
```

## Stack ในเครื่อง

```yaml title="compose.yaml"
services:
  api:
    build: .
    ports: ["8080:8080"]
    environment:
      DATABASE_URL: postgres://postgres:secret@db:5432/postgres?sslmode=disable
    depends_on: [db]
  db:
    image: postgres:17
    environment:
      POSTGRES_PASSWORD: secret
```

:::caution[ข้อควรระวัง]
- image แบบ static `scratch`/distroless ต้องใช้ `CGO_ENABLED=0` binary ที่ใช้ cgo จะพังด้วย "no such file or directory"
- ใส่ `.dockerignore` เพื่อไม่ให้ `.git`, `.env` และ binary ในเครื่องถูก copy เข้าไป
:::
