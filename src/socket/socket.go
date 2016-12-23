package main

import (
	l4g "base/log4go"
	"base/network"

	"os"
	"runtime"
)

var (
	g_server     = &SocketServer{}
	g_signal     = make(chan os.Signal, 1)
	g_socket_xml = new(XmlRecordConfigInfo)
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

	g_server = NewSocketServer()
	go network.TCPServe(g_server, conf)

	g_server.Run()
}
