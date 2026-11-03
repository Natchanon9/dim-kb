---
title: "Docker"
description: "A small, secure image with a multi-stage build"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Go compiles to one static binary, so the runtime image can be almost empty. A multi-stage Dockerfile builds in a Go image and copies only the binary.

## Dockerfile

```dockerfile title="Dockerfile"
FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download                 # cached unless dependencies change
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

## Local stack

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

:::caution[Pitfalls]
- `CGO_ENABLED=0` is required for `scratch`/distroless static images; a cgo binary fails with "no such file or directory".
- Add a `.dockerignore` so `.git`, `.env` and local binaries aren't copied in.
:::
