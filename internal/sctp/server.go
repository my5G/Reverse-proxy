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
		log.Fatalf("[PROXY] Error in Resolve Ip for reverse proxy")
	}

	// listen server
	go listenServer(ipAddr)

}

func listenServer(ipAddr *sctp.SCTPAddr) {

	ln, err := sctp.ListenSCTP("sctp", ipAddr)
	if err != nil {
		log.Fatalf("[GNB][PROXY] failed to listen: %v", err)
	}

	log.Printf("[GNB][PROXY] Listen on %s", ln.Addr())

	for {

		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to accept: %v", err)
		}

		log.Printf("[GNB][PROXY] Accepted Connection from RemoteAddr: %s", conn.RemoteAddr())

		wconn := conn.(*sctp.SCTPConn)

		wconn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)

		// handle connection with the remote client

		// select an AMF for the connection -- mapped AMF and GNB

		go serverClient(wconn)
	}
}

func serverClient(conn *sctp.SCTPConn) {

	for {

		buf := make([]byte, 65535+128)
		n, _, err := conn.SCTPRead(buf)
		if err != nil {
			log.Printf("[GNB] read failed: %v", err)
			return
		}

		// handle the read packets
		log.Printf("[GNB] read: %d packets", n)

		forwardData := make([]byte, n)
		copy(forwardData, buf[:n])
	}

}
