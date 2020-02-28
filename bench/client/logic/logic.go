package clientlogic

import (
	"fmt"
	"im/public/connect"
	"im/public/pb"
	"log"

	"github.com/golang/protobuf/proto"
)

//消息处理具柄
func ReceiveMsgFunc(connNode *connect.ConnNode, msgId int, pbMsg []byte) {
	content := &pb.MsgTestRep{}
	proto.Unmarshal(pbMsg, content)

	fmt.Printf("%v\n", content)

	// send msg
	p := &pb.MsgTestRep{
		MsgType:  int32(1),
		Username: string("client msg"),
		Age:      int32(33),
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode person:", err)
	}
	connNode.Encode(1, out)
}
