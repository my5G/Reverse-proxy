package main

import (
	"Reverse-proxy/internal/sctp"
	"time"
)

func main() {
	// sctp.InitServer("127.0.0.1", 38412)

	sctp.InitConn("127.0.0.2", 38412, "127.0.0.1", 38412)

	time.Sleep(1000 * time.Second)
}
