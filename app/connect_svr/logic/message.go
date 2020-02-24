package worldsvr_logic

import (
	"errors"
	"zh-im-go/public/msg"
)

type MsgFunc func(pbMsg []byte)
type SendPbMsg func(int, []byte)

var handler map[int]MsgFunc
var send SendPbMsg

func init() {
	handler = make(map[int]MsgFunc)
	send = nil

	registMsgHandler(msg.MSG_TEST_REQ, ConnSvrConnReq)
	registMsgHandler(msg.CONNSVR_CONN_RESP, ConnSvrConnResp)
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

func DealWithPbMsg(id int, msg []byte, cb func(int, []byte)) {
	if send == nil {
		send = cb
	}

	// dispatch msg
	successed := dispatch(id, msg)
	if !successed {
		//fmt.Println("Failed to get msg, for " + id)
	}
}

func SendPbMessage(msgType int, msg []byte) error {
	if send == nil {
		return errors.New("start or end postion is error")
	}

	send(msgType, msg)

	return nil
}
