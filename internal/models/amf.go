package models

import "github.com/ishidawataru/sctp"

const Inactive = 0x00
const Active = 0x01

type Amf struct {
	id      int
	Name    string `json:"name"`
	AmfIp   string `json:"ip"`
	AmfPort int    `json:"port"`
	State   int    `json:"state"`
	n2Amf   *sctp.SCTPConn
	n2Gnb   *sctp.SCTPConn
}

func (amf *Amf) getName() string {
	return amf.Name
}

func (amf *Amf) getId() int {
	return amf.id
}

func (amf *Amf) getAmfIp() string {
	return amf.AmfIp
}

func (amf *Amf) getAmfPort() int {
	return amf.AmfPort
}

func (amf *Amf) getN2Amf() *sctp.SCTPConn {
	return amf.n2Amf
}

func (amf *Amf) getN2Gnb() *sctp.SCTPConn {
	return amf.n2Gnb
}

func (amf *Amf) setN2Amf(conn *sctp.SCTPConn) {
	amf.n2Amf = conn
}

func (amf *Amf) setN2Gnb(conn *sctp.SCTPConn) {
	amf.n2Gnb = conn
}

func (amf *Amf) setInactiveAmf() {
	amf.State = Inactive
}

func (amf *Amf) setActiveAmf() {
	amf.State = Active
}
