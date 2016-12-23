package main
/*
import (
	l4g "base/log4go"
	"base/network"

	"protocol/out/cs"

	"github.com/golang/protobuf/proto"
)

func S2C_KeepAlive(sessioner network.TcpSessioner, ph *PackHead, buf []byte) bool {
	recv := &cs.S2C_KeepAlive{}
	err := proto.Unmarshal(buf, recv)
	if err != nil {
		l4g.Error("S2C_KeepAlive error: %v", err)
		return false
	}

	l4g.Debug("S2C_KeepAlive head: %+v, msg:%+v", ph, recv)

	return true
}

func S2C_Login(sessioner network.TcpSessioner, ph *PackHead, buf []byte) bool {
	recv := &cs.S2C_Login{}
	err := proto.Unmarshal(buf, recv)
	if err != nil {
		l4g.Error("S2C_Login error: %v", err)
		return false
	}

	l4g.Debug("S2C_Login head: %+v, msg:%+v", ph, recv)

	return true
}

func S2C_Flush(sessioner network.TcpSessioner, ph *PackHead, buf []byte) bool {
	recv := &cs.S2C_Flush{}
	err := proto.Unmarshal(buf, recv)
	if err != nil {
		l4g.Error("S2C_Flush error: %v", err)
		return false
	}

	l4g.Debug("S2C_Flush head: %+v, msg:%+v", ph, recv)

	return true
}
*/