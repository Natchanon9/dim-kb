---
title: "HTTP Client"
description: "Make outgoing requests with timeouts and context"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Create one `http.Client` with a timeout and reuse it; it keeps connections alive and is safe for concurrent use.

## Example

```go title="client.go"
client := &http.Client{Timeout: 5 * time.Second}

req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
req.Header.Set("Accept", "application/json")

res, err := client.Do(req)
if err != nil {
	return err // network error, timeout, cancelled ctx
}
defer res.Body.Close()
if res.StatusCode != http.StatusOK {
	return fmt.Errorf("get user: %s", res.Status) // 4xx/5xx is not an error from Do
}
var user User
err = json.NewDecoder(res.Body).Decode(&user)
```

## Checklist

- A timeout on the client and/or a context deadline.
- Always `defer res.Body.Close()`.
- Check `StatusCode` yourself.
- Reuse the client; don't create one per request.

:::caution[Pitfalls]
- `http.Get` and `http.DefaultClient` have **no timeout**; a hung server blocks forever.
- Not reading/closing the body prevents connection reuse.
:::
