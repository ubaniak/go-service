package main

import (
	"context"
	"go-service/internal/lib/config"
	"go-service/internal/server"
)

func main() {

	ctx := context.Background()
	s := server.New()

	config, err := config.Load(".")
	if err != nil {
		panic(err)
	}

	err = s.WithPort(config.Port).Run(ctx)

	// // err = s.WithPort(config.Port).WithMiddleware(handlers.Middleware()...).WithRoutes(handlers.Routes()...).Run(ctx)
	if err != nil {
		panic(err)
	}
}
