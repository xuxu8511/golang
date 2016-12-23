package main

import (
	l4g "base/log4go"
	"base/network"

	"io"
	"time"
)

type SocketSession struct {
	conn *network.TcpConn
	id   uint64
}

func NewSocketSession() *SocketSession {
	return &SocketSession{}
}

func (this *SocketSession) Init(conn *network.TcpConn) bool {
	this.conn = conn

	go this.Run()
	this.id = uint64(time.Now().UnixNano())
	g_server.sessions[this.id] = this

	return true
}

func (this *SocketSession) Process(buff []byte) bool {
	ph := &PackHead{}
	if ok := DecodePackHead(buff, ph); !ok {
		l4g.Error("decode pack head error")
		return false
	}

	l4g.Debug("local addr:%s, remote addr:%s, head: %+v, msg: %s", this.conn.LocalAddr(), this.conn.RemoteAddr(), ph, string(buff[PACK_HEAD_SIZE:]))

	this.WriteMsg(buff[PACK_HEAD_SIZE:])
	return true
}

func (this *SocketSession) Close() {
}

func (this *SocketSession) ReadMsg(r io.Reader) ([]byte, error) {
	var headBuf, bodyBuf []byte
	headBuf = make([]byte, PACK_HEAD_SIZE)
	if _, err := io.ReadFull(r, headBuf); err != nil {
		l4g.Error("read full fail: %+v", err)
		return nil, err
	}
	ph := &PackHead{}
	DecodePackHead(headBuf, ph)

	if ph.Len > PACK_HEAD_SIZE {
		msgLen := ph.Len - PACK_HEAD_SIZE
		bodyBuf = make([]byte, msgLen)
		if _, err := io.ReadFull(r, bodyBuf); err != nil {
			l4g.Error("read full fail: %+v", err)
			return nil, err
		}
	}

	msgBuf := make([]byte, 0, ph.Len)
	msgBuf = append(msgBuf, headBuf...)
	msgBuf = append(msgBuf, bodyBuf...)

	return msgBuf, nil
}

func (this *SocketSession) WriteMsg(msg []byte) bool {
	ph := &PackHead{
		Len: uint32(PACK_HEAD_SIZE + uint32(len(msg))),
		Cmd: 1,
	}
	buff := make([]byte, PACK_HEAD_SIZE)
	EncodePackHead(buff, ph)
	buff = append(buff, msg...)

	err := this.conn.AsyncWrite(buff, 0)
	if err != nil {
		l4g.Error("Write error: %+v", err)
		return false
	}
	return true
}

func (this *SocketSession) Run() {
	defer func() {
		delete(g_server.sessions, this.id)
		if err := recover(); err != nil {
			panic("Bug")
		}
	}()

	for {
		select {
		case msg := <-this.conn.PacketReceiveChan:
			this.Process(msg)
		case <-this.conn.CloseChan:
			l4g.Error("conn close chan")
			return
		}
	}
}
