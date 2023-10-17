package main

import (
	"posied0n/datacollection"

	"github.com/google/gopacket"
)

func main() {
	channel := make(chan gopacket.Packet) // Creating channel to write packets to.
	output := make(chan string)
	go datacollection.Startcapture(channel, "wlan0") // Calling Data Collector and passing the channel for writing packets to

	// // for p := range channel {
	// // 	apppayload := (p.ApplicationLayer()).Payload()
	// // 	if p.ApplicationLayer() != nil {
	// // 		fmt.Printf("%b\n",apppayload)
	// // 	}

	// // }
	// //go dataparser.Tcpcheck(channel, output)
	// // go dataparser.Ipcheck(channel,output)
	// // for {
	// fmt.Println(<-output)
	// // }
	// go dataparser.LayerCheck(channel, output)
	// for {
	// 	fmt.Println(<-output)
	// }
}
