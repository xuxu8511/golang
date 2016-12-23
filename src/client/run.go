package main

import (
	l4g "base/log4go"

	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	defer func() {
		if err := recover(); err != nil {
			l4g.Close()
			panic("Bug")
		}
	}()

	ticker := time.NewTicker(50000 * time.Millisecond)
	signal.Notify(g_signal, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case sig := <-g_signal:
			l4g.Debug("sig: %+v", sig)
			return
		case c := <-ticker.C:
			l4g.Debug("ticker: %+v", c)
			for i := 0; i < g_client_size; i++ {
				g_socket_session[i].WriteMsg([]byte("1234"))
			}
		}
	}
}
