---
title: "Graceful Shutdown"
description: "Finish in-flight requests before exiting"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

On SIGTERM (Kubernetes, Cloud Run, Docker stop), stop accepting new connections, let current requests finish, then exit. `http.Server.Shutdown` does that.

## Example

```go title="main.go"
srv := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second}

go func() {
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}()

ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()
<-ctx.Done() // wait for the signal

shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
if err := srv.Shutdown(shutdownCtx); err != nil {
	log.Println("forced shutdown:", err)
}
// then close DB pools, flush logs
```

`ListenAndServe` returns `http.ErrServerClosed` right away after `Shutdown`; that's expected, not a failure.

:::caution[Pitfalls]
- Keep the shutdown timeout below the platform's grace period (Kubernetes default 30s).
- `Shutdown` doesn't wait for hijacked connections or WebSockets; track those yourself.
:::
