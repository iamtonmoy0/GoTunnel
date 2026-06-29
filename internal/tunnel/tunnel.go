package tunnel

import (
	"io"
	"net"
)

func Copy(a net.Conn, b net.Conn) {
	go io.Copy(a, b)
	go io.Copy(b, a)
}
