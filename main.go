package main

import (
	"context"
	"go-service/internal/lib/config"
	"go-service/internal/server"
	"go-service/internal/settings"
)

func main() {

	ctx := context.Background()
	s := server.New()

	config, err := config.Load(".")
	if err != nil {
		panic(err)
	}

	middleware := settings.Middleware()
	routes := settings.Routes()

	err = s.WithPort(config.Port).WithMiddleware(middleware...).WithRoutes(routes...).Run(ctx)

	if err != nil {
		panic(err)
	}
}
