package networkinghttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

func Example_gracefulShutdown() {
	srv := &http.Server{Handler: http.NewServeMux(), ReadHeaderTimeout: 5 * time.Second}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	errCh := make(chan error, 1)
	go func() { errCh <- srv.Serve(ln) }()

	// in main this follows: ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM); <-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil { // stop accepting, wait for in-flight requests
		panic(err)
	}
	fmt.Println(errors.Is(<-errCh, http.ErrServerClosed))
	// Output: true
}
