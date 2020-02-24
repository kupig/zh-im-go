package main

import (
	"zh-im-go/public/config"
	"zh-im-go/public/connect"
)

func main() {
	tcpClient := connect.NewTCPClient(config.TCPServerAdrress)
	tcpClient.ClientStart(config.CONN_CLI)
}
