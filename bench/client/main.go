package main

import (
	clientlogic "im/bench/client/logic"
	"im/public/network"
	"reflect"
)

func main() {
	tcpClient := network.NewTCPClient("127.0.0.1:8899", reflect.ValueOf(clientlogic.ReceiveMsgFunc).Interface())
	tcpClient.Start()
}
