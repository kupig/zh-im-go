package worldsvr_logic

import (
	"zh-im-go/public/msg"
)

type MsgFunc func(pbMsg []byte)

var handler map[int]MsgFunc

func init() {
	handler = make(map[int]MsgFunc)

	registMsgHandler(msg.CONNSVR_CONN_REQ, ConnSvrConnReq)
}

func registMsgHandler(id int, cb MsgFunc) {
	_, ok := handler[id]
	if !ok {
		handler[id] = cb
	}
}

func dispatch(id int, msg []byte) bool {
	_, ok := handler[id]
	if !ok {
		cbfunc := handler[id]
		cbfunc(msg)

		return true
	}

	return false
}

func DealWithPbMsg(id int, msg []byte) {
	// for tests
	//content := &pb.MsgTestRep{}
	//proto.Unmarshal(pbMsg, content)
	//fmt.Println("world server message.")

	// dispatch msg
	successed := dispatch(id, msg)
	if !successed {
		//fmt.Println("Failed to get msg, for " + id)
	}
}
