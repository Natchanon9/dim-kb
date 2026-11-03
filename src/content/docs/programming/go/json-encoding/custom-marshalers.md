---
title: "Custom Marshalers"
description: "Control a type's JSON with MarshalJSON and UnmarshalJSON"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

A type that implements `json.Marshaler` / `json.Unmarshaler` decides its own JSON form. Use it for money, enums, custom date formats.

## Example: money in cents, JSON as "19.99"

```go title="money.go"
type Money int64 // cents

func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d.%02d"`, m/100, m%100)), nil
}

func (m *Money) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	var whole, frac int64
	if _, err := fmt.Sscanf(s, "%d.%d", &whole, &frac); err != nil {
		return fmt.Errorf("money %q: %w", s, err)
	}
	*m = Money(whole*100 + frac)
	return nil
}
// {"amount":1999 cents} ↔ {"amount":"19.99"}
```

For string-like values, `MarshalText`/`UnmarshalText` is simpler and also works for map keys and other encoders.

:::caution[Pitfalls]
- `UnmarshalJSON` needs a pointer receiver; `MarshalJSON` on a value receiver works for both values and pointers.
- Calling `json.Marshal(m)` inside `m.MarshalJSON` recurses forever; convert to the underlying type first.
:::
