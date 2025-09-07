package context

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	req := RequestInfo{
		Username: "Alice",
		IpAddr:   "192.168.1.1",
	}
	ctx = context.WithValue(ctx, "request", req)
	processRequest(ctx, req)
}
