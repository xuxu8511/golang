package main

import (
	l4g "base/log4go"
	"base/network"

	"io"
	"time"
)

type TemplateSession struct {
	id   uint64
	conn *network.TcpConn
}

func NewTemplateSession() *TemplateSession {
	return &TemplateSession{}
}

func (this *TemplateSession) Init(conn *network.TcpConn) bool {
	this.conn = conn

	go this.Run()
	this.id = uint64(time.Now().UnixNano())
	l4g.Debug("session init, local addr:%s, remote addr:%s", this.conn.LocalAddr(), this.conn.RemoteAddr())

	return true
}

func (this *TemplateSession) Process(buff []byte) bool {
	ph := &PackHead{}
	if ok := DecodePackHead(buff, ph); !ok {
		l4g.Error("decode pack head error")
		return false
	}

	g_CommandM.Dispatcher(this, ph, buff[PACK_HEAD_SIZE:])
	return true
}

func (this *TemplateSession) Close() {
}

func (this *TemplateSession) ReadMsg(r io.Reader) ([]byte, error) {
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
		if msgLen > this.conn.Conf.MaxReadMsgSize {
			l4g.Error("msg len overflow, len:%d, cmd:%d, uid:%d, sid:%d", ph.Len, ph.Cmd, ph.Uid, ph.Sid)
			return nil, network.ErrReadOverflow
		}

		bodyBuf = make([]byte, msgLen)
		if _, err := io.ReadFull(r, bodyBuf); err != nil {
			l4g.Error("read full fail: %+v", err)
			return nil, network.ErrReadBlocking
		}
	}

	msgBuf := make([]byte, 0, ph.Len)
	msgBuf = append(msgBuf, headBuf...)
	msgBuf = append(msgBuf, bodyBuf...)

	return msgBuf, nil
}

func (this *TemplateSession) WriteMsg(msg []byte) bool {
	ph := &PackHead{}
	ph.Len = uint32(PACK_HEAD_SIZE + uint32(len(msg)))
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

func (this *TemplateSession) Run() {
	defer func() {
		if err := recover(); err != nil {
			panic("Bug")
		}
	}()

	for {
		select {
		case msg := <-this.conn.PacketReceiveChan:
			this.Process(msg)
		case <-this.conn.CloseChan:
			return
		}
	}
}
