package app

import (
	"errors"
	"reflect"
)

type Container struct {
	registry map[string]interface{}
}

func New() *Container {
	return &Container{
		registry: make(map[string]interface{}),
	}
}

func (c *Container) Register(name string, i interface{}) error {
	_, ok := c.registry[name]
	if ok {
		return errors.New("dependency already registered")
	}
	c.registry[name] = i
	return nil
}

func (c *Container) Get(name string, i interface{}) error {
	s, ok := c.registry[name]
	if !ok {
		return errors.New("dependency not registered")
	}
	reflect.ValueOf(i).Elem().Set(reflect.ValueOf(s))
	return nil
}
