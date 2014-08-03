package client

import (
	"net/rpc"
	"regexp"
)

type Client rpc.Client

func Compile(s string) (*regexp.Regexp, error) {
	c, err := Connect()
	if err != nil {
		return nil, err
	}
	return c.Compile(s)
}

func Connect() (*Client, error) {
	c, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		return nil, err
	}
	return (*Client)(c), nil
}

func (c *Client) Compile(s string) (*regexp.Regexp, error) {
	var p string
	err := (*rpc.Client)(c).Call("Migemo.Pattern", s, &p)
	if err != nil {
		return nil, err
	}

	return regexp.Compile(p)
}
