package connect

import (
	"encoding/binary"
	"fmt"
	"net"
	connectsvr_logic "zh-im-go/app/connect_svr/logic"
	logicsvr_logic "zh-im-go/app/logic_svr/logic"
	worldsvr_logic "zh-im-go/app/world_svr/logic"
	"zh-im-go/public/config"
)

const (
	MsgTypeLen      = 2
	MsgBodyLen      = 2 // 消息内容长度
	ReadConnMaxLen  = 512
	WriteConnMaxLen = 512
)

type ConnNode struct {
	mngr    *ConnManager
	Conn    net.Conn
	rbuffer buffer
	wbuffer buffer
}

func (c *ConnNode) Read() (int, error) {
	return c.rbuffer.PushBuffer(c.Conn)
}

/*
func Encode() (int, []byte) {

}

func Decode() {

}
*/

func (c *ConnNode) Process(connCount, svrType int) error {
	defer func() {
		fmt.Printf("[关闭] %d TCP连接.\n", connCount)
		c.Release(connCount)
	}()

	fmt.Printf("[打开] %d 条TCP连接.\n", connCount)

	for {
		n, err := c.Read()
		if err != nil {
			return nil
		}

		if n > 0 {
			// MsgTypeLen
			typeLenBuf, err := c.rbuffer.GetBuffer(0, MsgTypeLen)
			if err != nil {
				continue
			}
			typeVal := int(binary.BigEndian.Uint16(typeLenBuf))
			fmt.Println(typeVal)

			// MsgBodyLen
			bodyLenBuf, err := c.rbuffer.GetBuffer(0, MsgBodyLen)
			if err != nil {
				continue
			}
			bodyLenVal := int(binary.BigEndian.Uint16(bodyLenBuf))
			fmt.Println(bodyLenVal)

			// MsgPbContent
			contentBuf, err := c.rbuffer.GetBuffer(0, bodyLenVal)
			if err != nil {
				continue
			}

			DistribudtionPbMsg(1, typeVal, contentBuf)
		}
	}
}

func (c *ConnNode) Release(connCount int) {
	c.Conn.Close()
	c.mngr.Release(c.rbuffer.buff, c.wbuffer.buff)
	fmt.Printf("[关闭] %d完全释放\n", connCount)
}

func DistribudtionPbMsg(svrType int, msgId int, pbMsg []byte) {
	switch svrType {
	case config.WORLD_SVR:
		worldsvr_logic.DealWithPbMsg(msgId, pbMsg)
	case config.CONN_SVR:
		connectsvr_logic.DealWithPbMsg(msgId, pbMsg)
	case config.LOGIC_SVR:
		logicsvr_logic.DealWithPbMsg(msgId, pbMsg)
	}
}
