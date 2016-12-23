package main

import (
	"base/network"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	l4g "base/log4go"
)

type SocketServer struct {
	sessions map[uint64]*SocketSession
}

func NewSocketServer() *SocketServer {
	return &SocketServer{
		sessions: make(map[uint64]*SocketSession),
	}
}

func (this *SocketServer) NewSession() network.TcpSessioner {
	return NewSocketSession()
}

func (this *SocketServer) Close() {
}

func (this *SocketServer) Run() {
	ticker := time.NewTicker(10000 * time.Millisecond)
	signal.Notify(g_signal, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case sig := <-g_signal:
			fmt.Println(sig)
			return
		case <-ticker.C:
			/*			for _, session := range this.sessions {
							session.Write([]byte("5678"))
						}
			*/
			l4g.Debug("server ticker")
		}
	}
}
