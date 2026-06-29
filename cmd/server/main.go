package main

import (
	"fmt"
	"github.com/iamtonmoy0/gotunnel/internal/tunnel"
	"net"
	"sync"
)

var (
	clients   []net.Conn
	clientsMu sync.Mutex
)

func main() {
	listener, err := net.Listen(
		"tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("GoTunnel server running on port :9000")
	fmt.Println("waiting for connections...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("new connection from:", conn.RemoteAddr())
		go handleConnection(conn)
	}

}

func handleConnection(user net.Conn) {
	defer user.Close()
	//
	client, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
		return
	}
	tunnel.Copy(
		user, client,
	)
}
