package connect

import (
	"encoding/binary"
	"errors"
	"fmt"
	"im/public/iobuffer"
	"im/public/message/msgconfig"
	msg "im/public/message/msgconfig"
	"net"
)

//链接节点
type ConnNode struct {
	//管理器
	manager *ConnManager
	//当前链接
	conn net.Conn
	//读buffer
	rbuffer iobuffer.Buffer
	//写buffer
	wbuffer iobuffer.Buffer
}

//读取bytes
func (c *ConnNode) Read() (int, error) {
	return c.rbuffer.PushBuffer(c.conn)
}

func (c *ConnNode) GetConnect() *net.Conn {
	return &c.conn
}

//消息编码
func (c *ConnNode) Decode() (int, []byte, error) {
	// MsgTypeLen
	byteMsgIdLen, err := c.rbuffer.GetBuffer(0, msgconfig.MsgTypeLen)
	if err != nil {
		return 0, nil, errors.New("Failed to call GetBuffer func.")
	}
	nMsgId := int(binary.BigEndian.Uint16(byteMsgIdLen))
	fmt.Println(nMsgId)

	// MsgBodyLen
	byteMsgBodyLen, err := c.rbuffer.GetBuffer(0, msgconfig.MsgBodyLen)
	if err != nil {
		return 0, nil, errors.New("Failed to call GetBuffer func.")
	}
	nBody := int(binary.BigEndian.Uint16(byteMsgBodyLen))
	fmt.Println(nBody)

	// MsgPbContent
	byteMsgContent, err := c.rbuffer.GetBuffer(0, nBody)
	if err != nil {
		return 0, nil, errors.New("Failed to call GetBuffer func.")
	}

	return nMsgId, byteMsgContent, nil
}

//消息解码
func (c *ConnNode) Encode(msgType int, bytes []byte) int {
	var buffer []byte
	if len(bytes) <= c.manager.WriteBufferMaxLen {
		cache := c.manager.WriteBuffer.Get().([]byte)
		buffer = cache[0 : msg.MsgTypeLen+msg.MsgBodyLen]

		defer c.manager.WriteBuffer.Put(cache)
	} else {
		// to do ...
	}

	// msg type
	binary.BigEndian.PutUint16(buffer[0:msgconfig.MsgTypeLen], uint16(msgType))

	// msg body len
	binary.BigEndian.PutUint16(buffer[msgconfig.MsgBodyLen:], uint16(len(bytes)))

	// msg content
	out := append(buffer, bytes...)

	// send msg
	n, err := c.conn.Write(out)
	if err != nil {
		fmt.Println(n)
	}

	return n
}
