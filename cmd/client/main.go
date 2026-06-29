package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to GoTunnel")
	conn.Write([]byte("Hello GoTunnel!"))
}
