---
title: "Regular Expressions"
description: "Match, extract and replace with regexp"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`regexp` uses RE2 syntax, which guarantees linear-time matching. Compile a pattern once and reuse it.

## Example

```go title="regexp.go"
var emailRe = regexp.MustCompile(`(\w+)@(\w+)\.com`) // compile once

emailRe.MatchString("a@b.com")                         // true
emailRe.FindStringSubmatch("mail bob@example.com now") // [bob@example.com bob example]
emailRe.ReplaceAllString("x@y.com", "$1 at $2")        // x at y
```

| Method | Returns |
| --- | --- |
| `MatchString` | Whether it matches |
| `FindString` / `FindAllString(s, -1)` | First / all matches |
| `FindStringSubmatch` | Match plus capture groups |
| `ReplaceAllString` | Replaced text (`$1` refers to groups) |

Use raw strings (backticks) for patterns so `\w` doesn't need escaping.

:::caution[Pitfalls]
- RE2 has no backreferences or lookahead.
- `MustCompile` panics on a bad pattern: fine for package-level constants, use `Compile` for user input.
- For simple checks, `strings.Contains`/`HasPrefix` is faster and clearer.
:::
