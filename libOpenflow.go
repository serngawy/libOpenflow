package main

import (
	"github.com/serngawy/libOpenflow/ofctrl"
)

func main() {

	// Main app test
	var app OfApp

	// Create a controller
	ctrler := ofctrl.NewController(&app)

	// start listening
	ctrler.Listen(":6633")
}