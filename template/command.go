package main

import (
	l4g "base/log4go"
	"base/network"
)

type Services func(network.TcpSessioner, *PackHead, []byte) bool

type CommandM struct {
	cmdm map[uint32]Services
}

func NewCommandM() *CommandM {
	return &CommandM{
		cmdm: make(map[uint32]Services),
	}
}

func (this *CommandM) Register(id uint32, service Services) {
	this.cmdm[id] = service
}

func (this *CommandM) Dispatcher(session network.TcpSessioner, ph *PackHead, data []byte) bool {
	if cmd, exist := this.cmdm[ph.Cmd]; exist {
		return cmd(session, ph, data)
	}
	l4g.Error("[Command] no find cmd: %d", ph.Cmd)
	return false
}

var g_CommandM = NewCommandM()

func InitCommand() {
}
