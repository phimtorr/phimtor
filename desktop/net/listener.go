package net

import (
	"errors"
	"net"
)

func CreateListener() (l net.Listener, listenPort int, close func() error, err error) {
	l, err = net.Listen("tcp", ":0")
	if err != nil {
		return nil, 0, nil, err
	}
	return l, l.Addr().(*net.TCPAddr).Port, func() error {
		err := l.Close()
		var opErr *net.OpError
		if errors.As(err, &opErr) && opErr.Op == "close" {
			return nil
		}
		return err
	}, nil
}
