---
title: "Publishing Modules"
description: "Tag a version and let the proxy and pkg.go.dev pick it up"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Publishing a Go module is pushing a git tag. There is no registry upload step.

## Steps

```sh
go mod tidy && go test ./...
git tag v1.2.0
git push origin v1.2.0
GOPROXY=proxy.golang.org go list -m github.com/me/lib@v1.2.0  # ask the proxy to fetch it
```

The module path in `go.mod` must match the repository URL. Docs appear on pkg.go.dev once the proxy has seen the version.

## Keeping it healthy

- Follow semver: breaking change → new major version and `/v2` path.
- Pulled a bad release? Add a `retract v1.2.0` directive in `go.mod` and publish a new version.
- Write doc comments on every exported name; they become the pkg.go.dev page.

:::caution[Pitfalls]
- Tags are effectively permanent: the proxy caches them. Never move or re-push a tag.
:::
