package main

import (
	"github.com/serngawy/libOpenflow/ofctrl"
	"fmt"
)

func main() {

	// Main app test
	var app OfApp

	// Create a controller
	ctrler := ofctrl.NewController(&app)

	// start listening
	fmt.Println("Starting OF controller at port 6633")
	ctrler.Listen(":6633")
}