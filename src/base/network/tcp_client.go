package network

import (
	"net"
)

func TCPClientServe(sessioner TcpSessioner, conf *Config) bool {
	conn, err := net.Dial("tcp", conf.Addr)
	if err != nil {
		return false
	}

	newTcpConn(conn, sessioner, conf).Do()
	return true
}
