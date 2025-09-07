package context

import (
	"context"
	"fmt"
	"time"
)

type RequestInfo struct {
	Username string
	IpAddr   string
}

func processRequest(ctx context.Context, request RequestInfo) {
	user := request.Username
	fmt.Println("Processing request from user %s at Ip %s\n", user, request.IpAddr)
	time.Sleep(2 * time.Second)
}
