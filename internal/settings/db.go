package settings

import (
	"errors"
	"go-service/internal/lib/dependency"
)

func RegisterDBs() (*dependency.Registry, error) {
	r := dependency.NewRegistry()
	return r, errors.New("not implimented")
}
