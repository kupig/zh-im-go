package serverlogic

import (
	"fmt"
	connect "im/public/connect"
	"im/public/pb"
	"log"

	"github.com/golang/protobuf/proto"
)

func ReceiveMsgFunc(connNode *connect.ConnNode, msgId int, pbMsg []byte) {
	content := &pb.MsgTestRep{}
	proto.Unmarshal(pbMsg, content)

	fmt.Printf("%v\n", content)

	// send msg
	p := &pb.MsgTestRep{
		MsgType:  int32(1),
		Username: string("to server msg"),
		Age:      int32(33),
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode person:", err)
	}
	connNode.Encode(1, out)
}
