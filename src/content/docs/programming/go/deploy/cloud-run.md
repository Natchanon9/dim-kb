---
title: "Cloud Run + Cloud SQL"
description: "Deploy a Go container to Cloud Run and connect to Cloud SQL Postgres"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Cloud Run runs a container per request load, scales to zero, and gives you HTTPS. The app only has to listen on `$PORT` and handle SIGTERM.

## App requirements

```go title="main.go"
port := os.Getenv("PORT") // set by Cloud Run, usually 8080
if port == "" {
	port = "8080"
}
srv := &http.Server{Addr: ":" + port, Handler: router}
// plus graceful shutdown on SIGTERM (see Networking & HTTP)
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

`--source .` builds with the Dockerfile (or buildpacks) via Cloud Build.

## Connecting to Cloud SQL

With `--add-cloudsql-instances`, Cloud Run exposes a Unix socket:

```text
host=/cloudsql/PROJECT:REGION:INSTANCE user=app password=... dbname=app
```

Alternatively use the Cloud SQL Go connector (`cloud.google.com/go/cloudsqlconn`) with IAM authentication.

:::caution[Pitfalls]
- Many instances × pool size can exceed Postgres `max_connections`; keep `SetMaxOpenConns` small or cap max instances.
- Secrets belong in Secret Manager (`--set-secrets`), not in plain env vars or the image.
:::
