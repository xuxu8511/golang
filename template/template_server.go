package main

import (
	l4g "base/log4go"
	"base/network"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type TemplateServer struct {
	sessions map[uint64]*TemplateSession
}

func NewTemplateServer() *TemplateServer {
	return &TemplateServer{
		sessions: make(map[uint64]*TemplateSession),
	}
}

func (this *TemplateServer) NewSession() network.TcpSessioner {
	return NewTemplateSession()
}

func (this *TemplateServer) Close() {
}

func (this *TemplateServer) Run() {
	ticker := time.NewTicker(10 * time.Second)
	signal.Notify(g_signal, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case sig := <-g_signal:
			l4g.Info("signal: %+v", sig)
			return
		case <-ticker.C:
			l4g.Debug("ticker...")
		}
	}
}
