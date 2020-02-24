package connect

import (
	"fmt"
	"log"
	"net"
	"zh-im-go/public/config"
	"zh-im-go/public/pb"

	"github.com/golang/protobuf/proto"
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

func (t *TCPClient) ClientStart(cliType int) {
	conn, err := net.Dial("tcp", config.TCPServerAdrress)
	if err != nil {
		fmt.Println("dial failed:", err)
	}
	defer conn.Close()

	connNode := cliConnManager.GetConnNode(conn)

	// send message
	go func() {
		p := &pb.MsgTestRep{
			MsgType:  int32(2),
			Username: string("helloworld"),
			Age:      int32(33),
		}

		out, err := proto.Marshal(p)
		if err != nil {
			log.Fatalln("Failed to encode person:", err)
		}
		connNode.Encode(2, out)
	}()

	// read message
	go func() {
		for {
			if connNode != nil {
				go connNode.CliProcess(cliType)
			}
		}
	}()

	//select {}
}
