package gowin32

import (
	"encoding/binary"
	"net"
	"unsafe"

	"github.com/winlabs/gowin32/wrappers"
)

// SendARP sends an ARP request. srcIP can be nil
func SendARP(destIP, srcIP net.IP) (net.HardwareAddr, error) {
	var s uint32
	if len(srcIP) > 0 {
		s = binary.LittleEndian.Uint32(srcIP.To4())
	}
	mac := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	macLen := uint32(len(mac))

	err := wrappers.SendARP(
		binary.LittleEndian.Uint32(destIP.To4()),
		s,
		(*uint32)(unsafe.Pointer(&mac[0])),
		&macLen)
	if err != nil {
		return nil, err
	}
	return mac, nil
}
