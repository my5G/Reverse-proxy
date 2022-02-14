package sctp

import (
	"Reverse-proxy/internal/models"
	"fmt"
	"github.com/ishidawataru/sctp"
	"log"
)

func InitServer(mgmt *models.Management) {

	address := fmt.Sprintf("%s:%d", mgmt.Ip, mgmt.Port)

	ipAddr, err := sctp.ResolveSCTPAddr("sctp", address)
	if err != nil {
		log.Fatalf("[PROXY] Error in Resolve Ip for reverse proxy")
	}

	// listen server
	go listenServer(ipAddr, mgmt)

}

func listenServer(ipAddr *sctp.SCTPAddr, mgmt *models.Management) {

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
		amf := mgmt.SelectAmfRb()

		go serverClient(wconn, amf)
	}
}

func serverClient(conn *sctp.SCTPConn, amf *models.Amf) {

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

		// send the data to the specific AMF
		amf.N2Amf.Write(forwardData)
	}
}
