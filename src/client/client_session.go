package main

import (
	l4g "base/log4go"
	"base/network"

	"io"
)

type ClientSession struct {
	conn *network.TcpConn
}

func NewClientSession() *ClientSession {
	return &ClientSession{}
}

func (this *ClientSession) Init(conn *network.TcpConn) bool {
	this.conn = conn

	go this.Receive()
	return true
}

func (this *ClientSession) Process(buff []byte) bool {
	ph := &PackHead{}
	if ok := DecodePackHead(buff, ph); !ok {
		l4g.Error("decode pack head error")
		return false
	}

	l4g.Debug("server return: local addr:%s, remote addr:%s, head: %+v, msg: %s", this.conn.LocalAddr(), this.conn.RemoteAddr(), ph, string(buff[PACK_HEAD_SIZE:]))

	return true
}

func (this *ClientSession) Close() {
}

func (this *ClientSession) ReadMsg(r io.Reader) ([]byte, error) {
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

func (this *ClientSession) WriteMsg(msg []byte) bool {
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

func (this *ClientSession) Receive() {
	for {
		select {
		case msg := <-this.conn.PacketReceiveChan:
			this.Process(msg)
		case <-this.conn.CloseChan:
			return
		}
	}
}
