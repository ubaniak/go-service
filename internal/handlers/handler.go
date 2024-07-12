package handlers

import (
	"go-service/internal/server"
)

func Middleware() []server.MiddlewareFunc {
	var m []server.MiddlewareFunc
	return m
}

func Routes() []server.RouteFunc {
	var r []server.RouteFunc = []server.RouteFunc{}
	return r
}
