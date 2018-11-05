## LibOpenflow and OF controller

Ofctrl implements a simple openflow controller using libOpenflow.

ofapp is the interface for Applcations "Consumers" in order to receiv events from the OF-Switch such as; connected, disconnected and packet-received.

# OF App implements the ConsumerInterface

    type OfApp struct {
      Switch *OFSwitch
    }

    func (o *OfApp) PacketRcvd(sw *OFSwitch, packet *openflow13.PacketIn) {
      log.Printf("App: Received packet: %+v", packet)
    }

    func (o *OfApp) SwitchConnected(sw *OFSwitch) {
      log.Printf("App: Switch connected: %v", sw.DPID())

      // Store switch for later use
      o.Switch = sw
    }

    func (o *OfApp) SwitchDisconnected(sw *OFSwitch) {
      log.Printf("App: Switch disconnected: %v", sw.DPID())
    }

    func (o *OfApp) MultipartReply(sw *OFSwitch, rep *openflow13.MultipartReply) {
      log.Println(rep.Body)
    }

# Example:

    func testExample() {

      // Main app
      var app ofctrl.OfApp

      // Create a controller
      ctrler := ofctrl.NewController(&app)

      // start listening
      ctrler.Listen(":6633")
    }

# Build:

We assume you already installed golang and dep. If not check the below links for more info
- install golang
 https://golang.org/doc/install
- install dep
 https://github.com/golang/dep/blob/master/docs/installation.md

To build the binary execute the ./build.sh script, it will create ofctrl binary under the $GOPATH bin directory.

Execute the ofctrl binary using the following command:
        $ sudo ./ofctrl

