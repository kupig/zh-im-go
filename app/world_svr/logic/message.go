package worldsvr_logic

import (
	"zh-im-go/public/msg"
)

type MsgFunc func(pbMsg []byte)

var handler map[int]MsgFunc

func init() {
	handler = make(map[int]MsgFunc)

	registMsgHandler(msg.MSG_TEST_REQ, ConnSvrConnReq)
	registMsgHandler(msg.CONNSVR_CONN_REQ, ConnSvrConnReq)
}

func registMsgHandler(id int, cb MsgFunc) {
	_, ok := handler[id]
	if !ok {
		handler[id] = cb
	}
}

func dispatch(id int, msg []byte) bool {
	cbfunc, ok := handler[id]
	if ok {
		//cbfunc := handler[id]
		cbfunc(msg)

		return true
	}

	return false
}

func DealWithPbMsg(id int, msg []byte) {
	// dispatch msg
	successed := dispatch(id, msg)
	if !successed {
		//fmt.Println("Failed to get msg, for " + id)
	}
}
