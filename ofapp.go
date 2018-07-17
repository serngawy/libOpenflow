package main

import(
	"github.com/serngawy/libOpenflow/openflow13"
	ofctrl "github.com/serngawy/libOpenflow/ofctrl"
	log "github.com/Sirupsen/logrus"
)

type OfApp struct {
	Switch *ofctrl.OFSwitch
}

func (app *OfApp) PacketRcvd(sw *ofctrl.OFSwitch, packet *openflow13.PacketIn) {
	log.Printf("App: Received packet: %+v", packet)
}

func (app *OfApp) SwitchConnected(sw *ofctrl.OFSwitch) {
	log.Printf("App: Switch connected: %v", sw.DPID())
	app.Switch = sw
	app.initPipline()
}

func (app *OfApp) SwitchDisconnected(sw *ofctrl.OFSwitch) {
	log.Printf("App: Switch disconnected: %v", sw.DPID())
}

func (app *OfApp) MultipartReply(sw *ofctrl.OFSwitch, rep *openflow13.MultipartReply) {
	log.Println(rep.Body)
}

func (app *OfApp) PortStatusChange(sw *ofctrl.OFSwitch, portStatus *openflow13.PortStatus) {
	log.Println("Port state: %+v", portStatus)
}

func (app *OfApp) FlowRemoved(sw *ofctrl.OFSwitch, flowRemoved *openflow13.FlowRemoved) {
	log.Println("Flow removed: %+v", flowRemoved)
}

//Here you define the App Pipeline tables
func (app *OfApp) initPipline() {
	//ex: set normal action on table 0
	flow := ofctrl.NewFlow(0)
	flow.SetNormalAction()
	log.Printf("App: flow key: %s", flow.FlowKey())
	app.Switch.InstallFlow(flow)

	// ex:match ip output port
	//flow := ofctrl.NewFlow(0)
	//ip := net.ParseIP("192.96.253.69")
	//flow.Match.IpDa = &ip
	//flow.FlowID = 100002
	//flow.Match.Ethertype = protocol.IPv4_MSG
	//flow.Match.IpProto = ofctrl.IP_PROTO_TCP
	//flow.Match.TcpDstPort = 8800
	//flow.Match.Priority = 100
	//flow.SetOutputPortAction(uint32(2))
	//flow.SetIPField(net.ParseIP("10.11.1.3"), "Dst")
	//
	//log.Printf("App: flow key: %s", flow.FlowKey())
	//app.Switch.InstallFlow(flow)
}