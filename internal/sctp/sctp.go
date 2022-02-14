package sctp

import (
	"Reverse-proxy/internal/models"
	"fmt"
	"github.com/ishidawataru/sctp"
	"log"
)

// code of sctp server
func InitServer(mgmt *models.Management) {

	address := fmt.Sprintf("%s:%d", mgmt.Ip, mgmt.PortServer)

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

		connGnb := conn.(*sctp.SCTPConn)

		connGnb.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)

		// handle connection with the remote client

		// select an amf for the connection, mapped amf and gnb
		// use round-robin
		amf := mgmt.SelectAmfRb()
		if amf == nil {
			log.Printf("[GNB][PROXY] Not AMF available")
		} else {
			// open SCTP connection with amf
			amf.N2Amf = InitConn(mgmt, amf)

			// increment the port of client
			mgmt.PortClient++

			// handle connections with amf
			go amfListen(connGnb, amf)

			// handle connections with proxy reverse
			go serverSctp(connGnb, amf)
		}
	}
}

func serverSctp(conn *sctp.SCTPConn, amf *models.Amf) {

	for {

		buf := make([]byte, 65535+128)
		n, _, err := conn.SCTPRead(buf)
		if err != nil {
			log.Printf("[GNB] read failed: %v", err)
			return
		}

		// handle the read packets
		log.Printf("[GNB][PROXY] read: %d packets", n)

		forwardData := make([]byte, n)
		copy(forwardData, buf[:n])

		// info about SCTP stream and PPID
		info := &sctp.SndRcvInfo{
			Stream: uint16(0),
			PPID:   0x3c000000,
		}

		// send the data to the specific AMF
		amf.N2Amf.SCTPWrite(forwardData, info)
	}
}

// code of connection with amf
func InitConn(mgmt *models.Management, amf *models.Amf) *sctp.SCTPConn {

	local := fmt.Sprintf("%s:%d", mgmt.Ip, mgmt.PortClient)
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

	return conn
}

func amfListen(conn *sctp.SCTPConn, amf *models.Amf) {

	for {

		buf := make([]byte, 65535)
		n, err := amf.N2Amf.Read(buf[:])
		if err != nil {
			log.Printf("[AMF] read failed: %v", err)
			return
		}

		// handle the read packets of AMF for the specific GNB
		log.Printf("[AMF][PROXY] read: %d packets", n)

		forwardData := make([]byte, n)
		copy(forwardData, buf[:n])

		// send the packets to GNB
		conn.Write(forwardData)
	}
}
