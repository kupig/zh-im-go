package main

import (
	serverlogic "im/bench/server/logic"
	"im/public/network"
	"reflect"
)

func main() {
	tcpServer := network.NewTCPServer("127.0.0.1:8899", reflect.ValueOf(serverlogic.ReceiveMsgFunc).Interface())
	tcpServer.Start()
}
