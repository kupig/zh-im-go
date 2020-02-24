package main

import (
	"zh-im-go/public/config"
	"zh-im-go/public/connect"
)

func main() {
	tcpClient := connect.NewTCPClient(config.TCPServerAdrress)
	tcpClient.ClientStart(config.WORLD_SVR)
	//time.Sleep(time.Second * 10)
	select {}
}
