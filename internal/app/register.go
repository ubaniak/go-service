package app

import (
	"go-service/internal/app/users"
	"go-service/internal/db"
)

func Register(db *db.ConnectionRegistry) *Container {
	c := NewContainer()
	c.Register(users.ServiceName, users.New("thin"))

	return c
}
