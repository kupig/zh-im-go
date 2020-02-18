package connect

type TCPClient struct {
	Address string
}

func NewTCPClient(address string) *TCPClient {
	return &TCPClient{
		Address: address,
	}
}

func (t *TCPClient) ClientStart() {
	/*
		conn, err := net.Dial("tcp", config.TCPServerAdrress)
		if err != nil {
			fmt.Println("dial failed:", err)
		}
		defer conn.Close()
	*/
}
