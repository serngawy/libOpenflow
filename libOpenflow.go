package libOpenflow

import (
	"github.com/serngawy/libOpenflow/ofctrl"
)

func testExample() {

	// Main app
	var app ofctrl.OfApp

	// Create a controller
	ctrler := ofctrl.NewController(&app)

	// start listening
	ctrler.Listen(":6633")
}