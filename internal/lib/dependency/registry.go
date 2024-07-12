package dependency

import (
	"errors"
	"reflect"
)

type Registry struct {
	registry map[string]interface{}
}

func NewRegistry() *Registry {
	return &Registry{
		registry: make(map[string]interface{}),
	}
}

func (c *Registry) Register(name string, i interface{}) error {
	_, ok := c.registry[name]
	if ok {
		return errors.New("dependency already registered")
	}
	c.registry[name] = i
	return nil
}

func (c *Registry) Get(name string, i interface{}) error {
	s, ok := c.registry[name]
	if !ok {
		return errors.New("dependency not registered")
	}
	reflect.ValueOf(i).Elem().Set(reflect.ValueOf(s))
	return nil
}
