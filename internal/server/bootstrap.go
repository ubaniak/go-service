package server

import (
	"context"
)

type server struct {
	port       string
	middleware []MiddlewareFunc
	routes     []RouteFunc
}

func New(port string) *server {
	return &server{
		port: port,
	}
}

func (s *server) WithMiddleware(m ...MiddlewareFunc) *server {
	s.middleware = m
	return s
}

func (s *server) WithRoutes(r ...RouteFunc) *server {
	s.routes = r
	return s
}

func (s *server) Bootstrap(ctx context.Context) error {
	r := setupRouter()

	addMiddleWare(ctx, r, s.middleware...)
	addRoutes(ctx, r, s.routes...)

	err := r.Run(":" + s.port)
	if err != nil {
		return err
	}

	return nil
}
