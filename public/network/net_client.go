package network

import (
	"errors"
	connect "im/public/connect"
	"im/public/message/msgconfig"
	"im/public/pb"
	"log"
	"net"
	"runtime"

	"github.com/golang/protobuf/proto"
)

type TCPClient struct {
	//服务器地址
	serverAddr string
	//消息回调函数
	callback interface{}
	//链接管理器
	connManager *connect.ConnManager
}

func NewTCPClient(address string, cb interface{}) *TCPClient {
	return &TCPClient{
		serverAddr:  address,
		callback:    cb,
		connManager: connect.CreateConnManager(msgconfig.MAX_READ_BUFFER_LEN, msgconfig.MAX_WRITE_BUFFER_LEN),
	}
}

func (client *TCPClient) Start() error {
	conn, err := net.Dial("tcp", client.serverAddr)
	if err != nil {
		return errors.New("Failed to get connect.")
	}

	defer func() {
		conn.Close()
	}()

	connNode := client.connManager.GetConnNode(conn)
	go client.SendMessage(connNode)
	go client.Process(connNode)

	for {
		runtime.Gosched()
	}
}

func (client *TCPClient) SendMessage(connNode *connect.ConnNode) {
	p := &pb.MsgTestRep{
		MsgType:  int32(1),
		Username: string("zhangyan123"),
		Age:      int32(33),
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode person:", err)
	}
	connNode.Encode(1, out)
}

func (client *TCPClient) Process(connNode *connect.ConnNode) error {
	for {
		_, err := connNode.Read()
		if err != nil {
			return errors.New("Failed to read buffer.")
		}

		nMsgId, byteCtx, err := connNode.Decode()
		if err != nil {
			return errors.New("Failed to decode buffer.")
		}

		if client.callback == nil {
			return errors.New("Failed to call message callback function.")
		}

		client.callback.(func(*connect.ConnNode, int, []byte))(connNode, nMsgId, byteCtx)
	}
}

func (client *TCPClient) Send(msgId int, ctx []byte) error {
	n := client.connManager.Connect.Encode(msgId, ctx)
	if n <= 0 {
		return errors.New("Failed to send message.")
	}

	return nil
}

func (client *TCPClient) Finish() {
	client.connManager.Release()
}
