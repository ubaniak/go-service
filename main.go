package main

import (
	"context"
	"go-service/internal/handlers"
	"go-service/internal/server"
)

func main() {

	ctx := context.Background()
	s := server.New()

	err := s.WithMiddleware(handlers.Middleware()...).WithRoutes(handlers.Routes()...).Run(ctx)
	if err != nil {
		panic(err)
	}
}
