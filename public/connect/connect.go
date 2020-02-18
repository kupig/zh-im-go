package connect

import (
	"encoding/binary"
	"fmt"
	"net"
	worldsvr_logic "zh-im-go/app/world_svr/logic"
	worldtsvr_logic "zh-im-go/app/world_svr/logc"
	logicsvr_logic "zh-im-go/app/logic_svr/logic"
)

const (
	MsgTypeLen      = 2
	MsgBodyLen      = 2 // 消息内容长度
	ReadConnMaxLen  = 512
	WriteConnMaxLen = 512
)

type ConnNode struct {
	mngr   *ConnManager
	Conn   net.Conn
	buffer buffer
}

func (c *ConnNode) Read() (int, error) {
	return c.buffer.PushBuffer(c.Conn)
}

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
			typeLenBuf, err := c.buffer.GetBuffer(0, MsgTypeLen)
			if err != nil {
				continue
			}
			typeVal := int(binary.BigEndian.Uint16(typeLenBuf))
			fmt.Println(typeVal)

			// MsgBodyLen
			bodyLenBuf, err := c.buffer.GetBuffer(0, MsgBodyLen)
			if err != nil {
				continue
			}
			bodyLenVal := int(binary.BigEndian.Uint16(bodyLenBuf))
			fmt.Println(bodyLenVal)

			// MsgPbContent
			contentBuf, err := c.buffer.GtBuffer(0, bodyLenVal)
			if err != nil {
				continue
			}

			DistribudtionPbMsg(1, typeVal, contentuf)
		}
}
}

fuc (c *ConnNode) Release(connCount int) {
	.Conn.Close()
c.mngr.Release(c.buffer.buff)
	fmt.Printf("[关闭] %d完全释放\n", connCount)
}

func DistribudtionPbMsg(svrType int, msId int, pbMsg []byte) {
	witch svrType {
case config.WORLD_SVR:
		worldsvr_logic.DealWithPbMsg(msgId, pbMsg)
	case config.CONNSVR:
		connectsvr_logic.DealithPbMsg(msgId, pbMsg)
	case config.LOGIC_SVR:
		logicsvr_logic.DealWthPbMsg(msgId, pbMsg)
	}
}
