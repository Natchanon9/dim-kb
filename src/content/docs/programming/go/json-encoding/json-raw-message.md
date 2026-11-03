---
title: "json.RawMessage"
description: "Delay decoding part of a JSON document"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`json.RawMessage` keeps a JSON fragment as raw bytes. Decode the envelope first, then decode the payload once you know its type.

## Example

```go title="event.go"
type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

var e Event
json.Unmarshal([]byte(`{"type":"user","data":{"name":"ann"}}`), &e)
switch e.Type {
case "user":
	var u struct {
		Name string `json:"name"`
	}
	json.Unmarshal(e.Data, &u) // ann
}
```

Typical uses: webhook and message-queue payloads with a `type` field, storing JSON columns without parsing them, passing JSON through unchanged.

:::caution[Pitfalls]
- A missing field leaves `RawMessage` nil, and `null` gives the bytes `null`; check both before decoding.
:::
