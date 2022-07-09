package mydb

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

// Driver mydb driver for implement database/sql/driver
type Driver struct {
}

func init() {
	log.Println("driver is call")
}

type Config struct {
	User   string // Username
	Passwd string // Password (requires User)
	Net    string // Network type
	DBName string // Database name
}

// Open for implement driver interface
func (driver *Driver) Open(name string) (driver.Conn, error) {
	log.Println("exec open driver")
	cfg, err := mysql.ParseDSN(name)
	fmt.Println(cfg)
	if err != nil {
		return nil, err
	}
	c := &connector{
		cfg: cfg,
	}
	return c.Connect(context.Background())
	//return &Conn{}, nil
}
