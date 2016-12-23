package network

import (
	l4g "base/log4go"

	"bufio"
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrConnClosing   = errors.New("use of closed network connection")
	ErrWriteBlocking = errors.New("write packet was blocking")
	ErrReadBlocking  = errors.New("read packet was blocking")
	ErrReadOverflow  = errors.New("read packet overflow")
)

type TcpSessioner interface {
	Init(*TcpConn) bool
	Process([]byte) bool
	ReadMsg(r io.Reader) ([]byte, error)
	WriteMsg(msg []byte) bool
	Close()
}

type TcpConn struct {
	conn      net.Conn
	sessioner TcpSessioner

	Conf *Config

	closeOnce         sync.Once
	closeFlag         int32
	CloseChan         chan struct{}
	PacketSendChan    chan []byte
	PacketReceiveChan chan []byte

	waitGroup *sync.WaitGroup
}

func newTcpConn(conn net.Conn, ser TcpSessioner, conf *Config) *TcpConn {
	return &TcpConn{
		conn:              conn,
		sessioner:         ser,
		Conf:              conf,
		CloseChan:         make(chan struct{}),
		PacketSendChan:    make(chan []byte, conf.WriteMsgQueueSize),
		PacketReceiveChan: make(chan []byte, conf.ReadMsgQueueSize),
		waitGroup:         &sync.WaitGroup{},
	}
}

func (this *TcpConn) Close() {
	if this.IsClosed() {
		return
	}

	this.closeOnce.Do(func() {
		atomic.StoreInt32(&this.closeFlag, 1)
		close(this.CloseChan)
		close(this.PacketSendChan)
		close(this.PacketReceiveChan)
		this.conn.Close()
		this.sessioner.Close()
	})
}

func (this *TcpConn) IsClosed() bool {
	return atomic.LoadInt32(&this.closeFlag) == 1
}

func (this *TcpConn) LocalAddr() string {
	return this.conn.LocalAddr().String()
}

func (this *TcpConn) RemoteAddr() string {
	return this.conn.RemoteAddr().String()
}

func (this *TcpConn) AsyncWrite(p []byte, timeout time.Duration) (err error) {
	if this.IsClosed() {
		return ErrConnClosing
	}

	defer func() {
		if e := recover(); e != nil {
			err = ErrConnClosing
		}
	}()

	if timeout == 0 {
		select {
		case this.PacketSendChan <- p:
			return nil
		case <-time.After(time.Duration(this.Conf.WriteTimeOut) * time.Second):
			return ErrWriteBlocking
		default:
			return ErrWriteBlocking
		}
	} else {
		select {
		case this.PacketSendChan <- p:
			return nil
		case <-this.CloseChan:
			return ErrConnClosing
		case <-time.After(timeout):
			return ErrWriteBlocking
		}
	}
}

func (this *TcpConn) Do() {
	if !this.sessioner.Init(this) {
		return
	}

	//	asyncDo(this.handleLoop, this.waitGroup)
	asyncDo(this.readLoop, this.waitGroup)
	asyncDo(this.writeLoop, this.waitGroup)
}

func (this *TcpConn) readLoop() {
	l4g.Debug("readLoop...")
	defer func() {
		recover()
		this.Close()
	}()

	reader := bufio.NewReader(this.conn)
	for {
		this.conn.SetReadDeadline(time.Now().Add(time.Duration(this.Conf.ReadTimeOut) * time.Second))
		p, err := this.sessioner.ReadMsg(reader)

		if err != nil {
			return
		}

		if READ_CHAN_LEN > 0 {
			select {
			case <-this.CloseChan:
				l4g.Debug("readloop close chan")
				return
			case this.PacketReceiveChan <- p:
			default:
				l4g.Debug("readloop default")
			}
		} else {
			if !this.sessioner.Process(p) {
				return
			}
		}
	}
}

func (this *TcpConn) writeLoop() {
	l4g.Debug("writeLoop...")
	defer func() {
		recover()
		this.Close()
	}()

	for {
		select {
		case <-this.CloseChan:
			return
		case p := <-this.PacketSendChan:
			if this.IsClosed() {
				return
			}
			if _, err := this.conn.Write(p); err != nil {
				return
			}
		}
	}
}

func (this *TcpConn) handleLoop() {
	l4g.Debug("handleLoop...")
	defer func() {
		recover()
		this.Close()
	}()

	for {
		select {
		case <-this.CloseChan:
			return
		case p := <-this.PacketReceiveChan:
			if this.IsClosed() {
				return
			}
			if !this.sessioner.Process(p) {
				return
			}
		}
	}
}

func asyncDo(fn func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		fn()
		wg.Done()
	}()
}
