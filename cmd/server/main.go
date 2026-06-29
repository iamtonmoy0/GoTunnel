package main

import (
	"fmt"
	"github.com/iamtonmoy0/gotunnel/internal/tunnel"
	"net"
)

func main() {
	listener, err := net.Listen(
		"tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("GoTunnel server running on port :9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("New tunnel connected")
		go handle(conn)
	}

}

func handle(user net.Conn) {
	defer user.Close()

	client, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
		return
	}
	tunnel.Copy(
		user, client,
	)
}
