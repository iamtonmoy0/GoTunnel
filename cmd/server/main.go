package main

import (
	"fmt"
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

func handle(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		fmt.Println(string(buffer[:n]))
	}

}
