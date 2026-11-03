---
title: "Module Versioning"
description: "Semantic versions, upgrades and the /v2 rule"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Modules use semantic versions `vMAJOR.MINOR.PATCH`. Go picks the **minimum** version that satisfies every requirement (MVS), so builds don't change when a new release appears.

## Choosing versions

```sh
go get github.com/foo/bar@v1.4.2  # exact version
go get github.com/foo/bar@latest  # newest release
go get -u ./...                    # upgrade all direct and indirect deps (minor/patch)
go get github.com/foo/bar@none    # remove
```

## Major versions

From v2 onward, the major version is part of the import path: `github.com/foo/bar/v2`. v1 and v2 can coexist in one build.

## Other terms

- **Pseudo-version** (`v0.0.0-20240101120000-abcdef123456`): a commit without a tag.
- **`replace`** in `go.mod`: point a dependency at a fork or local folder.
- **`// indirect`**: needed by a dependency, not by your code directly.

:::caution[Pitfalls]
- `v0.x` promises no compatibility; expect breaking changes on minor bumps.
- Don't commit `replace` lines pointing to your local disk.
:::
