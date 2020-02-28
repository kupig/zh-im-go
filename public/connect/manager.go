package connect

import (
	"im/public/iobuffer"
	"net"
	"sync"
)

type ConnManager struct {
	//当前链接信息
	Connect *ConnNode
	//最大读长度
	ReadBufferMaxLen int
	//最大写长度
	WriteBufferMaxLen int
	//读buffer
	ReadBuffer sync.Pool
	//写buffer
	WriteBuffer sync.Pool
}

// 创建链接管理器
func CreateConnManager(rbufferMaxLen, wbufferMaxLen int) *ConnManager {
	return &ConnManager{
		Connect:           nil,
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
	mngr.Connect = &ConnNode{
		manager: mngr,
		conn:    c,
		rbuffer: iobuffer.NewBuffer(mngr.ReadBuffer.Get().([]byte)),
		wbuffer: iobuffer.NewBuffer(mngr.WriteBuffer.Get().([]byte)),
	}
	return mngr.Connect
}

func (mngr *ConnManager) Release() {
	mngr.Connect.conn.Close()

	mngr.ReadBuffer.Put(mngr.Connect.rbuffer)
	mngr.WriteBuffer.Put(mngr.Connect.rbuffer)
}
