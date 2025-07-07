package main

import (
	"context"

	"github.com/guiestimo/bank-simulator-hexagonal/internal/adapter/input/controller/startup"
	"github.com/guiestimo/bank-simulator-hexagonal/internal/config"
)

func main() {
	ctx := context.Background()

	config.Parse(ctx)

	server := startup.NewServer()
	server.Start()
}
