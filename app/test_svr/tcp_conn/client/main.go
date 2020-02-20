package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
	"zh-im-go/public/config"
	MSG "zh-im-go/public/msg"
	"zh-im-go/public/pb"

	"github.com/golang/protobuf/proto"
)

func send(c net.Conn) {

	p := &pb.MsgTestRep{
		MsgType:  int32(1),
		Username: string("helloworld"),
		Age:      int32(33),
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode person:", err)
	}

	pbLen := len(out)
	sendBuf := make([]byte, MSG.MsgTypeLen+MSG.MsgBodyLen)

	// MsgTypeLen
	binary.BigEndian.PutUint16(sendBuf, uint16(MSG.MSG_TEST_REQ))

	// MsgBodyLen
	binary.BigEndian.PutUint16(sendBuf[2:], uint16(pbLen))
	fmt.Println(uint16(pbLen))

	// MsgPbContent
	buf := append(sendBuf, out...)
	c.Write(buf)
}

func main() {

	buffer := make([]byte, 4)
	binary.BigEndian.PutUint16(buffer, uint16(257))

	conn, err := net.Dial("tcp", config.TCPServerAdrress)
	if err != nil {
		fmt.Println("dial failed:", err)
	}
	defer conn.Close()
	send(conn)

	time.Sleep(time.Second * 10)

	//tcpClient := connect.NewTCPClient(config.TCPServerAdrress)
	//tcpClient.ClientStart()
}
