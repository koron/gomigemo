package server

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/koron/gomigemo/embedict"
)

func RunDefault() error {
	d, err := embedict.Load()
	if err != nil {
		return err
	}
	migemo := &Migemo{dict: d}

	rpc.Register(migemo)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		return err
	}
	http.Serve(l, nil)

	return nil
}
