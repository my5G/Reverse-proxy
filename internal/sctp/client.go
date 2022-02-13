package sctp

import (
	"fmt"
	"github.com/ishidawataru/sctp"
	"log"
)

func InitConn(ipLocal string, portLocal int, ipRemote string, portRemote int) {

	local := fmt.Sprintf("%s:%d", ipLocal, portLocal)
	remote := fmt.Sprintf("%s:%d", ipRemote, portRemote)

	localAddr, err := sctp.ResolveSCTPAddr("sctp", local)
	if err != nil {
		log.Fatalf("[AMF] Error in Resolve Ip Local for reverse proxy")
	}

	remoteAddr, err := sctp.ResolveSCTPAddr("sctp", remote)
	if err != nil {
		log.Fatalf("[AMF] Error in Resolve Ip Remote for reverse proxy")
	}

	conn, err := sctp.DialSCTPExt(
		"sctp",
		localAddr,
		remoteAddr,
		sctp.InitMsg{NumOstreams: 2, MaxInstreams: 2})
	if err != nil {
		log.Fatalf("[AMF] Error in connect to AMF for reverse proxy")
	}

	log.Println("[AMF] AMF connected")

	conn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)

	go amfListen(conn)
}

func amfListen(conn *sctp.SCTPConn) {

	for {

		buf := make([]byte, 65535)
		n, _, err := conn.SCTPRead(buf[:])
		if err != nil {
			log.Printf("[AMF] read failed: %v", err)
			return
		}

		// handle the read packets of AMF for the specific GNB
		log.Printf("[AMF] read: %d packets", n)

		forwardData := make([]byte, n)
		copy(forwardData, buf[:n])
	}
}
