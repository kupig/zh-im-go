package connect

import (
	"fmt"
	"net"
	"time"
	"zh-im-go/public/config"
)

var cliConnManager = CreateConnManager(ReadConnMaxLen, WriteConnMaxLen)

type TCPClient struct {
	Address string
}

func NewTCPClient(address string) *TCPClient {
	return &TCPClient{
		Address: address,
	}
}

func (t *TCPClient) ClientStart(svrType int) {
	conn, err := net.Dial("tcp", config.TCPServerAdrress)
	if err != nil {
		fmt.Println("dial failed:", err)
	}
	defer conn.Close()

	for {
		connNode := cliConnManager.GetConnNode(conn)
		if connNode != nil {
			connNode.CliProcess(svrType)
		}

		time.Sleep(time.Second * 1)
	}
}
