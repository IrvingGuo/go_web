package main

import (
	"context"
	"go_web/server"
)

func main() {
	server.NewServer(context.Background())
}
