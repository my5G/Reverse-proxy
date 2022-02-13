package sctp

import (
	"fmt"
	"github.com/ishidawataru/sctp"
	"log"
)

func InitServer(ip string, port int) {

	address := fmt.Sprintf("%s:%d", ip, port)

	ipAddr, err := sctp.ResolveSCTPAddr("sctp", address)
	if err != nil {
		log.Fatalf("Error in Resolve Ip for reverse proxy")
	}

	// listen server
	go listenServer(ipAddr)

}

func listenServer(ipAddr *sctp.SCTPAddr) {

	fmt.Println("1")

	ln, err := sctp.ListenSCTP("sctp", ipAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listen on %s", ln.Addr())

	for {

		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to accept: %v", err)
		}

		log.Printf("Accepted Connection from RemoteAddr: %s", conn.RemoteAddr())

		wconn := conn.(*sctp.SCTPConn)

		wconn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)

		// handle connection with the remote client
		go serverClient(wconn)
	}
}

func serverClient(conn *sctp.SCTPConn) {

	for {

		buf := make([]byte, 65535+128)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("read failed: %v", err)
			return
		}

		// handle the read packets
		log.Printf("read: %d packets", n)
	}

}
