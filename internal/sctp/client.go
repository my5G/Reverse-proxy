package sctp

import (
	"Reverse-proxy/internal/models"
	"fmt"
	"github.com/ishidawataru/sctp"
	"log"
)

func InitConn(mgmt *models.Management, amf *models.Amf) *sctp.SCTPConn {

	local := fmt.Sprintf("%s:%d", mgmt.Ip, mgmt.Port)
	remote := fmt.Sprintf("%s:%d", amf.AmfIp, amf.AmfPort)

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

	go amfListen(conn, amf)

	return conn
}

func amfListen(conn *sctp.SCTPConn, amf *models.Amf) {

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

		// send the packets to GNB
		amf.N2Gnb.Write(forwardData)
	}
}
