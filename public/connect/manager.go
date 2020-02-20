package connect

import (
	"net"
	"sync"
	//"zh-im-go/public/connec"
	//"zh-im-go/public/connect"
)

type ConnManager struct {
	ReadBufferMaxLen  int
	WriteBufferMaxLen int
	ReadBuffer        sync.Pool
	WriteBuffer       sync.Pool
}

func CreateConnManager(rbufferMaxLen, wbufferMaxLen int) *ConnManager {
	return &ConnManager{
		ReadBufferMaxLen:  rbufferMaxLen,
		WriteBufferMaxLen: wbufferMaxLen,
		ReadBuffer: sync.Pool{
			New: func() interface{} {
				b := make([]byte, rbufferMaxLen)
				return b
			},
		},
		WriteBuffer: sync.Pool{
			New: func() interface{} {
				b := make([]byte, wbufferMaxLen)
				return b
			},
		},
	}
}

func (mngr *ConnManager) GetConnNode(c net.Conn) *ConnNode {
	return &ConnNode{
		mngr:    mngr,
		Conn:    c,
		rbuffer: NewBuffer(mngr.ReadBuffer.Get().([]byte)),
		wbuffer: NewBuffer(mngr.WriteBuffer.Get().([]byte)),
	}
}

func (mngr *ConnManager) Release(rbytes, wbytes []byte) {
	mngr.ReadBuffer.Put(rbytes[0:0])
	mngr.WriteBuffer.Put(wbytes[0:0])
}
