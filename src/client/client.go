package main

import (
	l4g "base/log4go"
	"base/network"

	"os"
	"runtime"
)

var (
	g_client_size    = 100
	g_socket_session = make([]*ClientSession, g_client_size)
	g_signal         = make(chan os.Signal, 1)
	g_socket_xml     = new(XmlClientConfigInfo)
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if !LoadConf() {
		l4g.Error("load config failed")
		return
	}

	l4g.LoadConfiguration(g_socket_xml.Log.Config)
	defer l4g.Close()

	conf := &network.Config{
		Addr: g_socket_xml.Server.Addr,
	}
	l4g.Debug("conf: %+v", conf)

	for i := 0; i < g_client_size; i++ {
		g_socket_session[i] = NewClientSession()
		if !network.TCPClientServe(g_socket_session[i], conf) {
			l4g.Error("tcp connect fail socket server: %v", conf)
			return
		}
	}

	Run()
}
