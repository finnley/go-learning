package mydb

import (
	"database/sql/driver"
	"errors"
	"github.com/go-sql-driver/mysql"
	"net"
)

// Conn for db open
type Conn struct {
	buf     buffer
	netConn net.Conn
	cfg     *mysql.Config
}

// Prepare statement for prepare exec
func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return &MyStmt{}, nil
}

// Close close db connection
func (c *Conn) Close() error {
	return errors.New("can't close connection")
}

// Begin begin
func (c *Conn) Begin() (driver.Tx, error) {
	return nil, errors.New("not support tx")
}
