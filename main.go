package main

import (
	"context"
	"go-service/internal/app"
	"go-service/internal/db"
	"go-service/internal/handlers"
	"go-service/internal/server"
)

func main() {

	ctx := context.Background()
	s := server.New("8080")

	db, err := db.Register()
	if err != nil {
		panic(err)
	}
	c := app.Register(db)

	err = s.WithMiddleware(handlers.Middleware()...).WithRoutes(handlers.Routes(c)...).Bootstrap(ctx)
	if err != nil {
		panic(err)
	}
}
