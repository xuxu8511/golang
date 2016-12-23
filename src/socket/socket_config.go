package main

import (
	"base/common"
	l4g "base/log4go"
)

type RecordConfigLog struct {
	Config string `xml:"config"`
}

type RecordConfigServer struct {
	Addr string `xml:"port"`
}

type XmlRecordConfigInfo struct {
	Log    RecordConfigLog    `xml:"log"`
	Server RecordConfigServer `xml:"server"`
}

func LoadConf() bool {
	dir := "../config/"
	filename := ""

	filename = dir + "socket_config.xml"
	g_socket_xml = new(XmlRecordConfigInfo)
	if err := common.LoadConfig(filename, g_socket_xml); err != nil {
		l4g.Error("load config %s failed: %v", filename, err)
		return false
	}
	return true
}
