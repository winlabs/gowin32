package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/winlabs/gowin32"
)

var (
	// ipFlag is used to set an IPv4 address destination for an ARP request
	dstIPFlag = flag.String("d", "", "IPv4 address destination for ARP request")

	// ifaceFlag is used to set a network interface for ARP requests
	srcIPFlag = flag.String("s", "", "IPv4 source address of the sender")
)

func main() {
	flag.Parse()

	if len(*dstIPFlag) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	mac, err := gowin32.SendARP(net.ParseIP(*dstIPFlag).To4(), net.ParseIP(*srcIPFlag).To4())
	if err != nil {
		fmt.Printf("could not find MAC for %s: %v", *dstIPFlag, err)
		os.Exit(1)
	}
	fmt.Printf("MAC address for %s is %v\n", *dstIPFlag, mac)
}
