package dataparser

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func Tcpcheck(c chan gopacket.Packet, op chan string) {
	for {
		for i := range c {
			if tcplayer := (i.Layer(layers.LayerTypeTCP)); tcplayer != nil {
				tcp, _ := tcplayer.(*layers.TCP)
				op <- fmt.Sprintf("%d srcport , %d dstport", tcp.SrcPort, tcp.DstPort)
			}
		}
	}
}
func Ipcheck(c chan gopacket.Packet, op chan string) {
	for {
		for i := range c {
			if iplayer := (i.Layer(layers.LayerTypeIPv4)); iplayer != nil {
				ip, _ := iplayer.(*layers.IPv4)
				op <- fmt.Sprintf("%d srcIP , %d dstIP", ip.SrcIP, ip.DstIP)
			}
		}

	}
}
