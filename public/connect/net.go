package connect

import (
	"fmt"
	"net"
	"os"
	"zh-im-go/public/config"
)

var connManager = CreateConnManager(ReadConnMaxLen)

type TCPServer struct {
	Address string
	Port    string
}

type TCPClient struct {
	Address string
}

func NewTCPServer(address string) *TCPServer {
	return &TCPServer{
		Address: address,
	}
}

func (t *TCPServer) Start(svrType int) {
	addr, err := net.ResolveTCPAddr("tcp", t.Address)
	if err != nil {
		fmt.Printf("error: ", err)
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Printf("error: ", err)
		os.Exit(1)
	}

	var connCount int = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Printf("error: ", err)
			os.Exit(1)
		}

		connNode := connManager.GetConnNode(conn)
		if connNode != nil {
			go connNode.Process(connCount, svrType)
			connCount++
		}
	}
}


func NewTCPClient(address string) *TCPClient{
	return &TCPClient{
		Address: address,
	}
}

func (t *TCPClient) ClientStart() {
	conn, err := net.Dial("tcp", config.TCPServerAdrress)
	if err != nil {
		fmt.Println("dial failed:", err)
	}
	defer conn.Close()
}

