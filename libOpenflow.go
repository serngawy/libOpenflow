package main

import (
	"github.com/serngawy/libOpenflow/ofctrl"
)

func main() {

	// Main app test
	var app ofctrl.OfApp

	// Create a controller
	ctrler := ofctrl.NewController(&app)

	// start listening
	ctrler.Listen(":6633")
}