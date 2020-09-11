package gowin32

import (
	"encoding/binary"
	"fmt"
	"net"
	"unsafe"

	"github.com/winlabs/gowin32/wrappers"
)

// SendARP sends an ARP request. srcIP can be nil
func SendARP(destIP, srcIP net.IP) (net.HardwareAddr, error) {
	var s uint32
	if len(srcIP) > 0 {
		srcIPv4 := srcIP.To4()
		if srcIPv4 == nil {
			return nil, fmt.Errorf("%s is not valid IPv4 address", srcIP)
		}
		s = binary.LittleEndian.Uint32(srcIPv4)
	}
	destIPv4 := destIP.To4()
	if destIPv4 == nil {
		return nil, fmt.Errorf("%s is not valid IPv4 address", destIP)
	}
	mac := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	macLen := uint32(len(mac))

	err := wrappers.SendARP(
		binary.LittleEndian.Uint32(destIPv4),
		s,
		(*uint32)(unsafe.Pointer(&mac[0])),
		&macLen)
	if err != nil {
		return nil, err
	}
	return mac, nil
}
