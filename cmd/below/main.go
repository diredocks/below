package main

import (
	"context"
	"fmt"

	"below/internal/comment"
	"below/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("Hello, Below!")
	comment.InitDB()
	app := server.NewServer(ctx)
	app.Listen(":3000")
}
