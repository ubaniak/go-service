package main

import (
	"context"
	"go-service/internal/app"
	"go-service/internal/app/users"
	"go-service/internal/handlers"
	"go-service/internal/server"
)

func main() {

	ctx := context.Background()
	s := server.New("8080")

	c := app.New()
	c.Register(users.ServiceName, users.New("test"))

	err := s.WithMiddleware(handlers.Middleware()...).WithRoutes(handlers.Routers(c)...).Bootstrap(ctx)
	if err != nil {
		panic(err)
	}
}
