package main

import (
	l4g "base/log4go"
	"base/network"

	"os"
	"runtime"
)

var (
	g_socket_xml = new(XmlTemplateConfigInfo)
	g_signal     = make(chan os.Signal, 1)
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if !LoadConf() {
		l4g.Error("load config failed")
		return
	}

	l4g.LoadConfiguration(g_socket_xml.Log.Config)
	defer l4g.Close()

	InitCommand()

	conf := &network.Config{
		Addr:              g_socket_xml.Server.Addr,
		MaxReadMsgSize:    g_socket_xml.Server.MaxReadMsgSize,
		ReadMsgQueueSize:  g_socket_xml.Server.ReadMsgQueueSize,
		ReadTimeOut:       g_socket_xml.Server.ReadTimeOut,
		MaxWriteMsgSize:   g_socket_xml.Server.MaxWriteMsgSize,
		WriteMsgQueueSize: g_socket_xml.Server.WriteMsgQueueSize,
		WriteTimeOut:      g_socket_xml.Server.WriteTimeOut,
	}
	l4g.Debug("conf: %+v", conf)
}
