package db

import (
	"fmt"

	"gorm.io/gorm"
)

type ConnectionRegistry struct {
	registry map[string]*gorm.DB
}

func NewConnectionRegistry() *ConnectionRegistry {
	return &ConnectionRegistry{registry: make(map[string]*gorm.DB)}
}

func (d *ConnectionRegistry) Register(name string, db *gorm.DB) error {
	_, ok := d.registry[name]
	if ok {
		return fmt.Errorf("database %s is already registered", name)
	}

	d.registry[name] = db
	return nil
}

func (d *ConnectionRegistry) Get(name string) (*gorm.DB, error) {
	db, ok := d.registry[name]
	if !ok {
		return nil, fmt.Errorf("database %s is not registered", name)
	}
	return db, nil
}
