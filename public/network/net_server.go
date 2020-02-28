package network

import (
	"errors"
	"fmt"
	connect "im/public/connect"
	"im/public/message/msgconfig"
	"net"
	"os"
)

type TCPServer struct {
	//服务器的地址
	addr string
	//消息回掉函数
	callback interface{}
	//链接管理器
	connManager *connect.ConnManager
}

//创建一个链接
func NewTCPServer(address string, cb interface{}) *TCPServer {
	return &TCPServer{
		addr:        address,
		callback:    cb,
		connManager: connect.CreateConnManager(msgconfig.MAX_READ_BUFFER_LEN, msgconfig.MAX_WRITE_BUFFER_LEN),
	}
}

func (server *TCPServer) Start() {
	addr, err := net.ResolveTCPAddr("tcp", server.addr)
	if err != nil {
		fmt.Printf("error: ", err)
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Printf("error: ", err)
		os.Exit(1)
	}
	defer func() {
		listener.Close()
		server.Finish()
	}()

	var connCount int = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Printf("error: ", err)
			os.Exit(1)
		}

		connNode := server.connManager.GetConnNode(conn)
		if connNode != nil {
			go server.Process(connNode)
			connCount++
		}
	}
}

func (server *TCPServer) Process(c *connect.ConnNode) {
	for {
		n, err := c.Read()
		if err != nil {
			continue
		}

		if n > 0 {
			nMsgId, byteCtx, err := c.Decode()
			if err != nil {
				continue
			}
			//fmt.Println(nMsgId, byteCtx)
			if server.callback == nil {
				break
			}
			server.callback.(func(*connect.ConnNode, int, []byte))(c, nMsgId, byteCtx)
		}
	}
}

func (server *TCPServer) Send(msgId int, ctx []byte) error {
	n := server.connManager.Connect.Encode(msgId, ctx)
	if n <= 0 {
		return errors.New("Failed to send message.")
	}

	return nil
}

func (server *TCPServer) Finish() {
	server.connManager.Release()
}
