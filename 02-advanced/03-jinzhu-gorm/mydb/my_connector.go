package mydb

import (
	"bytes"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net"
	"strings"
)

type connector struct {
	cfg *mysql.Config // immutable private copy.
}

func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {
	var err error

	log.Println("exec open connector")
	mc := &Conn{
		cfg: c.cfg,
	}

	nd := net.Dialer{Timeout: mc.cfg.Timeout}
	mc.netConn, err = nd.DialContext(ctx, mc.cfg.Net, mc.cfg.Addr)

	mc.buf = newBuffer(mc.netConn)

	//fmt.Fprintf(mc.netConn, "GET / HTTP/1.0\r\n\r\n")

	//buf := make([]byte, 0, 4096) // big buffer
	//tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	//for {
	//	n, err := mc.netConn.Read(tmp)
	//	if err != nil {
	//		if err != io.EOF {
	//			fmt.Println("read error:", err)
	//		}
	//		break
	//	}
	//	//fmt.Println("got", n, "bytes.")
	//	buf = append(buf, tmp[:n]...)
	//
	//}
	//fmt.Println("total size:", len(buf))
	//log.Println(string(buf))
	//fmt.Fprintf(mc.netConn, "GET / HTTP/1.0\r\n\r\n")
	var buf bytes.Buffer
	_, err = io.Copy(&buf, mc.netConn)
	if err != nil {
		return nil, err
	}
	fmt.Println("total size:", buf.Len())
	log.Println(buf.String())

	if strings.Contains(buf.String(), "mysql") {
		//err := errors.New(buf.String())
		mc.netConn.Close()
		mc.netConn = nil
		return nil, errors.New(buf.String())
	}

	return mc, err
}

func (c *connector) Driver() driver.Driver {
	return &Driver{}
}

type buffer struct {
	buf []byte
	nc  net.Conn
	cfg *Config
}

func newBuffer(nc net.Conn) buffer {
	fg := make([]byte, 4096)
	return buffer{
		buf: fg,
		nc:  nc,
	}
}
