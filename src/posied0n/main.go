package main

import (
	"fmt"
	"posied0n/datacollection"

	"github.com/google/gopacket"
)

func main() {
	channel := make(chan gopacket.Packet)   // Creating channel to write packets to.
	go datacollection.Startcapture(channel,"wlan0") // Calling Data Collector and passing the channel for writing packets to

	for p := range channel {
		fmt.Println(p.ApplicationLayer())
		fmt.Printf("%b\n",(p.ApplicationLayer()).Payload())
	}
}
