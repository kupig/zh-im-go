package msg

// 消息头部信息
// msg = head + body
// head = msgTypeLen + msgBodyLen
// body = pbBuff
const (
	MsgTypeLen = 2 // 消息类型长度
	MsgBodyLen = 2 // 消息内容长度
)

// 服务器链接消息
const (
	//MSG_TEST_REQ = 1
	CONNSVR_CONN_REQ	= 1
	CONNSVR_CONN_RESP	 = 2
)

type MsgBase interface {
	virtualTypeFun() int
	logic() interface{}
}

type MsgTestReq struct {
	msgType  int32
	userName string
	age      int32
}

