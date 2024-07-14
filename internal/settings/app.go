package settings

import (
	"errors"
	"go-service/internal/lib/dependency"
)

func RegisterApps(dbs *dependency.Registry) (*dependency.Registry, error) {
	r := dependency.NewRegistry()
	return r, errors.New("not implimented")
}
