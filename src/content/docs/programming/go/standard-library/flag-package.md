---
title: "flag Package"
description: "Parse command-line flags"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`flag` parses `-name=value` arguments into typed variables, with defaults and generated `-h` help.

## Example

```go title="main.go"
port := flag.Int("port", 8080, "listen port")
debug := flag.Bool("debug", false, "enable debug logs")
flag.Parse()
fmt.Println(*port, *debug, flag.Args())
// ./app -port=9000 -debug extra → 9000 true [extra]
```

Other forms: `flag.StringVar(&cfg.Addr, "addr", ":8080", "...")`, `flag.Duration`, and `flag.NewFlagSet` for subcommands.

## Flags vs env vars

Services in containers usually read config from environment variables; flags suit CLIs and local runs. Many apps accept both, with env as the default value of the flag.

:::caution[Pitfalls]
- Flags must come before positional arguments: `app -v file`, not `app file -v`.
- For complex CLIs with subcommands, `spf13/cobra` is the common choice.
:::
