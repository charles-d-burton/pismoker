package main

import (
	"flag"
	"log"

	"github.com/charles-d-burton/pismoker/controller"
)

var usageStr = `
Usage: pismoker [options]
Options:
	-nh, --nats-host <NATSHost>     Start the controller connecting to the defined NATS Streaming server
`

func usage() {
	log.Fatalf(usageStr)
}

// NOTE: Use tls scheme for TLS, e.g. stan-sub -s tls://demo.nats.io:4443 foo
func main() {
	var natsHost string

	flag.StringVar(&natsHost, "nh", "", "Start the controller connecting to the defined NATS Streaming server")
	flag.StringVar(&natsHost, "nats-host", "", "Start the controller connecting to the defined NATS Streaming server")
	flag.Parse()
	if natsHost == "" {
		usage()
	}
	controller.StartServer(natsHost)
}
