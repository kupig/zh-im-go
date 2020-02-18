package msg

import (
	"zh-im-go/public/msg"
)

type MsgFunc func(msg.MsgBase)

var handler map[int]MsgFunc

func init() {
	handler = make(map[int]MsgFunc)
	//registMsgHandler()
}

func registMsgHandler(e int, cb MsgFunc) {
	_, ok := handler[e]
	if !ok {
		handler[e] = cb
	}
}

func FooMsgTestReq(msg.MsgBase) {
}
