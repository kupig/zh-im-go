package worldsvr_logic

import (
	"fmt"
	"zh-im-go/public/pb"

	"github.com/golang/protobuf/proto"
)

func ConnSvrConnReq(pbMsg []byte) {
	content := &pb.MsgTestRep{}
	proto.Unmarshal(pbMsg, content)
	fmt.Println("world server message.")
	fmt.Println("new connect server will connected.")
}

func ConnSvrConnResp() {

}
