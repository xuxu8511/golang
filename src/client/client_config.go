package main

import (
	"base/common"
	l4g "base/log4go"
)

type ClientConfigLog struct {
	Config string `xml:"config"`
}

type ClientConfigServer struct {
	Addr string `xml:"port"`
}

type XmlClientConfigInfo struct {
	Log    ClientConfigLog    `xml:"log"`
	Server ClientConfigServer `xml:"server"`
}

func LoadConf() bool {
	dir := "../config/"
	filename := ""

	filename = dir + "client_config.xml"
	g_socket_xml = new(XmlClientConfigInfo)
	if err := common.LoadConfig(filename, g_socket_xml); err != nil {
		l4g.Error("load config %s failed: %v", filename, err)
		return false
	}

	return true
}
