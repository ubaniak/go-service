package main

import (
	"context"
	"go-service/internal/app"
	"go-service/internal/handlers"
	"go-service/internal/server"
)

func main() {

	ctx := context.Background()
	s := server.New("8080")

	c := app.Register()

	err := s.WithMiddleware(handlers.Middleware()...).WithRoutes(handlers.Routes(c)...).Bootstrap(ctx)
	if err != nil {
		panic(err)
	}
}
