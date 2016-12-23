package network

import (
	"net"
)

type Server interface {
	NewSession() TcpSessioner
	Close()
}

func TCPServe(srv Server, conf *Config) {
	defer srv.Close()

	l, e := net.Listen("tcp", conf.Addr)
	if e != nil {
		panic(e.Error())
	}
	defer l.Close()

	for {
		conn, e := l.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				continue
			}
			return
		}
		newTcpConn(conn, srv.NewSession(), conf).Do()
	}
}
