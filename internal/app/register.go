package app

import "go-service/internal/app/users"

func Register() *Container {
	c := NewContainer()
	c.Register(users.ServiceName, users.New("thin"))

	return c
}
