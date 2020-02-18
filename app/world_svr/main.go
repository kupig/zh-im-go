package main

import (
	"zh-im-go/public/config"
	"zh-im-go/public/connect"
)

func main() {
	tcpServer := connect.NewTCPServer(config.TCPServerAdrress)
	tcpServer.Start(config.WORLD_SVR)
}
