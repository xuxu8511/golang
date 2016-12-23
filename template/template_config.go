package main

import (
	"base/common"
	l4g "base/log4go"
)

type TemplateConfigLog struct {
	Config string `xml:"config"`
}

type TemplateConfigServer struct {
	Addr              string `xml:"port"`
	MaxReadMsgSize    uint32 `xml:"MaxReadMsgSize"`
	ReadMsgQueueSize  uint32 `xml:"ReadMsgQueueSize"`
	ReadTimeOut       uint32 `xml:"ReadTimeOut"`
	MaxWriteMsgSize   uint32 `xml:"MaxWriteMsgSize"`
	WriteMsgQueueSize uint32 `xml:"WriteMsgQueueSize"`
	WriteTimeOut      uint32 `xml:"WriteTimeOut"`
}

type XmlTemplateConfigInfo struct {
	Log    TemplateConfigLog    `xml:"log"`
	Server TemplateConfigServer `xml:"server"`
}

func LoadConf() bool {
	dir := "../config/"
	filename := ""

	filename = dir + "template_config.xml"
	g_socket_xml = new(XmlTemplateConfigInfo)
	if err := common.LoadConfig(filename, g_socket_xml); err != nil {
		l4g.Error("load config %s failed: %v", filename, err)
		return false
	}
	return true
}
