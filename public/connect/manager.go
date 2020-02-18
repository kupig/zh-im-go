package connect

import (
	"net"
	"sync"
	//"zh-im-go/public/connec"
	//"zh-im-go/public/connect"
)

type ConnManager struct {
	ReadBufferMaxLen int
	ReadBuffer       sync.Pool
}

func CreateConnManager(readBufferMaxLen int) *ConnManager {
	return &ConnManager{
		ReadBufferMaxLen: readBufferMaxLen,
		ReadBuffer: sync.Pool{
			New: func() interface{} {
				b := make([]byte, readBufferMaxLen)
				return b
			},
		},
	}
}

func (mngr *ConnManager) GetConnNode(c net.Conn) *ConnNode {
	return &ConnNode{
		mngr:   mngr,
		Conn:   c,
		buffer: NewBuffer(mngr.ReadBuffer.Get().([]byte)),
	}
}

func (mngr *ConnManager) Release(bytes []byte) {
	bytes = bytes[0:0]
	mngr.ReadBuffer.Put(bytes)
}
