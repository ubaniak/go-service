package handlers

import (
	"go-service/internal/app"
	"go-service/internal/app/users"
	"go-service/internal/server"
)

func Middleware() []server.MiddlewareFunc {
	var m []server.MiddlewareFunc
	return m
}

func Routers(c *app.Container) []server.RouteFunc {
	var r []server.RouteFunc = []server.RouteFunc{}
	var u users.UserService
	err := c.Get(users.ServiceName, &u)
	if err != nil {
		panic(err)
	}

	r = append(r, users.Routes(u)...)
	return r
}
