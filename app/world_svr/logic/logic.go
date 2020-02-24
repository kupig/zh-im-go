package worldsvr_logic

import (
	"fmt"
	"log"
	"zh-im-go/public/msg"
	"zh-im-go/public/pb"

	"github.com/golang/protobuf/proto"
)

func ConnSvrConnReq(pbMsg []byte) {
	content := &pb.MsgTestRep{}
	proto.Unmarshal(pbMsg, content)
	fmt.Println("world server message.")

	// send test
	p := &pb.MsgTestRep{
		MsgType:  int32(3),
		Username: string("helloworld"),
		Age:      int32(33),
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode person:", err)
	}
	SendPbMessage(msg.CONNSVR_CONN_RESP, out)

}

func ConnSvrConnResp() {

}
