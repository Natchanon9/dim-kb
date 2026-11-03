---
title: "time Package"
description: "Times, durations, formatting and parsing"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`time.Time` is an instant with a location; `time.Duration` is nanoseconds as an `int64`. Formatting uses a **reference date** instead of `%Y-%m-%d`.

## Example

```go title="time.go"
t := time.Date(2026, 11, 3, 14, 30, 0, 0, time.UTC)
t.Format("2006-01-02 15:04")            // 2026-11-03 14:30
t.Format(time.RFC3339)                  // 2026-11-03T14:30:00Z
t.Weekday()                             // Tuesday
t.Add(90 * time.Minute).Format(time.Kitchen) // 4:00PM

d, _ := time.ParseDuration("1h30m")     // d.Minutes() → 90
xmas, _ := time.Parse("02/01/2006", "25/12/2026")
xmas.Sub(t).Hours()                     // 1233.5
```

The reference date is `Mon Jan 2 15:04:05 MST 2006` (1 2 3 4 5 6 7).

## Habits

- Store and transmit times in UTC (RFC 3339); convert to local only for display (`time.LoadLocation("Asia/Bangkok")`).
- Compare with `t.Before`, `t.After`, `t.Equal`, not `==`.
- Measure elapsed time with `time.Since(start)`.

:::caution[Pitfalls]
- `time.Parse` without a zone gives UTC; use `ParseInLocation` for local input.
- `time.Sleep` in tests makes them slow and flaky.
:::
