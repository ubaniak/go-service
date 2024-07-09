package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type connectionBuilder struct {
	driver   string
	host     string
	port     string
	user     string
	password string
	database string
}

func New() *connectionBuilder {
	return &connectionBuilder{}
}

func (c *connectionBuilder) SetDriver(d string) *connectionBuilder {
	c.driver = d
	return c
}

func (c *connectionBuilder) SetHost(h string) *connectionBuilder {
	c.host = h
	return c
}

func (c *connectionBuilder) SetPassword(p string) *connectionBuilder {
	c.password = p
	return c
}

func (c *connectionBuilder) SetUser(u string) *connectionBuilder {
	c.user = u
	return c
}

func (c *connectionBuilder) SetDatabase(d string) *connectionBuilder {
	c.database = d
	return c
}

func (c *connectionBuilder) Build() (*gorm.DB, error) {
	var dsn string
	var dialector gorm.Dialector
	switch c.driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.user, c.password, c.host, c.port, c.database)
		dialector = mysql.Open(dsn)
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.host, c.port, c.user, c.password, c.database)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", c.driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
