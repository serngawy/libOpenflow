package ofctrl

import(
	"github.com/serngawy/libOpenflow/openflow13"
	log "github.com/Sirupsen/logrus"
)

type OfApp struct {
	Switch *OFSwitch
}

func (o *OfApp) PacketRcvd(sw *OFSwitch, packet *openflow13.PacketIn) {
	log.Printf("App: Received packet: %+v", packet)
}

func (o *OfApp) SwitchConnected(sw *OFSwitch) {
	log.Printf("App: Switch connected: %v", sw.DPID())
	o.Switch = sw

}

func (o *OfApp) SwitchDisconnected(sw *OFSwitch) {
	log.Printf("App: Switch disconnected: %v", sw.DPID())
}

func (o *OfApp) MultipartReply(sw *OFSwitch, rep *openflow13.MultipartReply) {
	log.Println(rep.Body)
}

//Here you define the App Pipeline tables
func (o *OfApp) initPipline() {
	//ex: set normal action on table 0
	flow := NewFlow(0)
	flow.SetNormalAction()
	log.Printf("App: flow key: %s", flow.flowKey())
	o.Switch.InstallFlow(flow)

	//ex: set drop action for vlan 49 in table 0
	flow = NewFlow(0)
	flow.Match.VlanId = 49
	flow.Match.Priority = 100
	flow.SetDropAction()
	log.Printf("App: flow key: %s", flow.flowKey())
	o.Switch.InstallFlow(flow)
}