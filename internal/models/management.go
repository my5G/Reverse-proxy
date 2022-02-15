package models

type Management struct {
	Ip         string
	PortClient int
	PortServer int
	amfList    [20]*Amf
	round      int
	amfLength  int
}

func InitMgmt(ip string, portClient int, portServer int) *Management {
	mgmt := &Management{}
	mgmt.round = 0
	mgmt.amfLength = 0
	mgmt.Ip = ip
	mgmt.PortClient = portClient
	mgmt.PortServer = portServer
	return mgmt
}

func (mgmt *Management) CreateAmf(amf *Amf) {
	mgmt.amfList[mgmt.amfLength] = amf
	mgmt.amfLength++
}

func (mgmt *Management) UpdateAmfState(amf Amf, name string) (*Amf, bool) {

	for i := 0; i < mgmt.amfLength; i++ {
		if mgmt.amfList[i].Name == name {
			mgmt.amfList[i].State = amf.State
			return mgmt.amfList[i], true
		}
	}
	return nil, false
}

func (mgmt *Management) SelectAmfRb() *Amf {
	if mgmt.amfLength == 0 {
		return nil
	} else {
		for i := 0; i < mgmt.amfLength; i++ {
			if mgmt.amfList[mgmt.round].State == 0 {
				mgmt.updateRoundRobin()
			} else {
				index := mgmt.round
				mgmt.updateRoundRobin()
				return mgmt.amfList[index]
			}
		}
	}
	return nil
}

func (mgmt *Management) updateRoundRobin() {
	mgmt.round++
	mgmt.round = mgmt.round % (mgmt.amfLength)
}

func (mgmt *Management) UpdateMgmtPort() {
	mgmt.PortClient++
}
