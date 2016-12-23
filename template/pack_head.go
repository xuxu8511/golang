package main

import (
	l4g "base/log4go"
	"base/network"

	"errors"
	"io"
)

const (
	PACK_HEAD_SIZE uint32 = 16
)

type PackHead struct {
	Len uint32
	Cmd uint32
	Uid uint32
	Sid uint32
}

func DecodePackHead(buf []byte, ph *PackHead) bool {
	if len(buf) < int(PACK_HEAD_SIZE) {
		l4g.Error("decode error")
		return false
	}
	ph.Len = network.DecodeUint32(buf[0:])
	ph.Cmd = network.DecodeUint32(buf[4:])
	ph.Uid = network.DecodeUint32(buf[8:])
	ph.Sid = network.DecodeUint32(buf[12:])
	return true
}

func EncodePackHead(buf []byte, ph *PackHead) bool {
	if len(buf) < int(PACK_HEAD_SIZE) {
		l4g.Error("encode error")
		return false
	}
	network.EncodeUint32(ph.Len, buf[0:])
	network.EncodeUint32(ph.Cmd, buf[4:])
	network.EncodeUint32(ph.Uid, buf[8:])
	network.EncodeUint32(ph.Sid, buf[12:])
	return true
}
