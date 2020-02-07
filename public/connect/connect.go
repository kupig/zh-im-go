package connect

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	//"zh-im-go/public/msg"
	"zh-im-go/public/pb"
)

const (
	MsgTypeLen      = 2
	MsgBodyLen = 2 // 消息内容长度
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
			content := &pb.MsgTestRep{}
			contentBuf, err := c.buffer.GetBuffer(0, bodyLenVal)
			if err != nil {
				continue
			}
			proto.Unmarshal(contentBuf, content)
			fmt.Println(bodyLenVal)
		}
	}
}

func (c *ConnNode) Release(connCount int) {
	c.Conn.Close()
	c.mngr.Release(c.buffer.buff)
	fmt.Printf("[关闭] %d完全释放\n", connCount)
}
