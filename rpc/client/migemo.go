package client

import (
	"net/rpc"
	"regexp"
)

// Client is a migemo RPC client
type Client rpc.Client

// Compile compiles a migemo string and get regexp.Regexp with migemo RPC
// server.
func Compile(s string) (*regexp.Regexp, error) {
	c, err := Connect()
	if err != nil {
		return nil, err
	}
	return c.Compile(s)
}

// Addr is migemo RPC server address.
var Addr = "127.0.0.1:1234"

// Connect connects migemo RPC server and get Client.
func Connect() (*Client, error) {
	c, err := rpc.DialHTTP("tcp", Addr)
	if err != nil {
		return nil, err
	}
	return (*Client)(c), nil
}

// Compile compiles a migemo string and get regexp.Regexp.
func (c *Client) Compile(s string) (*regexp.Regexp, error) {
	var p string
	err := (*rpc.Client)(c).Call("Migemo.Pattern", s, &p)
	if err != nil {
		return nil, err
	}

	return regexp.Compile(p)
}
