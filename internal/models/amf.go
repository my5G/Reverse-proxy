package models

import "github.com/ishidawataru/sctp"

const Inactive = 0x00
const Active = 0x01

type Amf struct {
	Name    string `json:"name"`
	AmfIp   string `json:"ip"`
	AmfPort int    `json:"port"`
	State   int    `json:"state"`
	N2Amf   *sctp.SCTPConn
}

func (amf *Amf) getName() string {
	return amf.Name
}

func (amf *Amf) getAmfIp() string {
	return amf.AmfIp
}

func (amf *Amf) getAmfPort() int {
	return amf.AmfPort
}

func (amf *Amf) setInactiveAmf() {
	amf.State = Inactive
}

func (amf *Amf) setActiveAmf() {
	amf.State = Active
}
