package main

import (
	"fmt"
)

type IPAddr [4]byte

func bytetostr(b byte) string {
	if (int(b) > 10) {
		return bytetostr(b/10) + string((int(b)%10) + 0x30)
	} else {
		return string((int(b)%10) + 0x30)
	}
}

func (ip IPAddr) String() string {
	out := ""
	out += bytetostr(ip[0]) + "."
	out += bytetostr(ip[1]) + "."
	out += bytetostr(ip[2]) + "."
	out += bytetostr(ip[3]) 
	return out
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}