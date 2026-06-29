package main

import (
	"github.com/iamtonmoy0/gotunnel/internal/tunnel"
	"net"
)

func main() {
	server, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	local, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
	tunnel.Copy(
		server, local)
}
