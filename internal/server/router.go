package server

import (
	"context"

	"github.com/gin-gonic/gin"
)

type RouteFunc func(context.Context, *gin.Engine)
type MiddlewareFunc func(context.Context) gin.HandlerFunc

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func addMiddleWare(ctx context.Context, router *gin.Engine, middleware ...MiddlewareFunc) {
	for _, m := range middleware {
		router.Use(m(ctx))
	}
}

func addRoutes(ctx context.Context, router *gin.Engine, routes ...RouteFunc) {
	for _, fn := range routes {
		fn(ctx, router)
	}
}
