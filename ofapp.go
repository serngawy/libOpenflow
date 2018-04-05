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

//Here you define the App Pipeline tables
func (app *OfApp) initPipline() {
	//ex: set normal action on table 0
	flow := ofctrl.NewFlow(0)
	flow.SetNormalAction()
	log.Printf("App: flow key: %s", flow.FlowKey())
	app.Switch.InstallFlow(flow)
}