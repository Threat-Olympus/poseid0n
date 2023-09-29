package datacollection

import (
	//"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

func Startcapture(v chan gopacket.Packet,itf string ) {
	phyInterface := itf // Interface selection to be made dynamic
	handle, err := pcap.OpenLive(phyInterface, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for {
		v <- <-packetSource.Packets()
	}

}
