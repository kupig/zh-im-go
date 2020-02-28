package main

import (
	"time"
	"zh-im-go/public/config"
	"zh-im-go/public/connect"
)

func main() {
	tcpClient := connect.NewTCPClient(config.TCPServerAdrress)
	tcpClient.Start(config.WORLD_SVR)
	time.Sleep(time.Second * 10)
}
