package gocontext

import (
	"context"
	"fmt"
)

type requestIDKey struct{} // unexported key type: no collisions with other packages

func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, id)
}

func requestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey{}).(string)
	return id, ok
}

func Example_withValue() {
	ctx := withRequestID(context.Background(), "req-42")
	fmt.Println(requestID(ctx))
	fmt.Println(requestID(context.Background()))
	// Output:
	// req-42 true
	//  false
}
