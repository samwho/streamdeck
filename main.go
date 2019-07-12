package main

import (
	"flag"
)

func main() {
	var port int
	var pluginUUID string
	var registerEvent string
	var info string

	flag.IntVar(&port, "name", -1, "")
	flag.StringVar(&pluginUUID, "pluginUUID", "", "")
	flag.StringVar(&registerEvent, "registerEvent", "", "")
	flag.StringVar(&info, "info", "", "")
	flag.Parse()

}
