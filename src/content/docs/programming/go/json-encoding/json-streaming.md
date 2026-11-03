---
title: "JSON Streaming"
description: "Decode and encode JSON directly from readers and writers"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`json.NewDecoder(r)` reads JSON from any `io.Reader`; `json.NewEncoder(w)` writes to any `io.Writer`. Use them for HTTP bodies, files and streams of values.

## Example

```go title="stream.go"
dec := json.NewDecoder(strings.NewReader(`{"id":1} {"id":2}`))
for {
	var e struct{ ID int }
	err := dec.Decode(&e)
	if errors.Is(err, io.EOF) {
		break
	}
	if err != nil {
		return err
	}
	fmt.Println(e.ID) // 1, then 2
}

json.NewEncoder(os.Stdout).Encode(map[string]string{"status": "ok"})
// {"status":"ok"} plus a newline
```

## In HTTP handlers

```go
dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
dec.DisallowUnknownFields()
if err := dec.Decode(&req); err != nil { ... }
```

:::caution[Pitfalls]
- `Decode` reads one value; trailing garbage after it isn't an error unless you check `dec.More()` or decode again.
:::
