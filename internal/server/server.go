package server

import (
	"context"
)

type server struct {
	port       string
	middleware []MiddlewareFunc
	routes     []RouteFunc
}

func New() *server {
	return &server{
		port: "8080",
	}
}

func (s *server) WithPort(p string) *server {
	s.port = p
	return s
}

func (s *server) WithMiddleware(m ...MiddlewareFunc) *server {
	s.middleware = m
	return s
}

func (s *server) WithRoutes(r ...RouteFunc) *server {
	s.routes = r
	return s
}

func (s *server) Run(ctx context.Context) error {
	r := setupRouter()

	addMiddleWare(ctx, r, s.middleware...)
	addRoutes(ctx, r, s.routes...)

	err := r.Run(":" + s.port)
	if err != nil {
		return err
	}

	return nil
}
