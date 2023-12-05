package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"poseid0n/ruleengine"
)

func main() {
	// Read rules from a file
	rules, err := readRulesFromFile("rules/example.rules")
	if err != nil {
		log.Fatal("Error reading rules:", err)
	}

	// Create a RuleEngine with the provided rules
	engine := ruleengine.NewRuleEngine(rules)

	// Open the network interface for packet capture
	handle, err := pcap.OpenLive("en0", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Process packets from the interface and match against rules
	for packet := range packetSource.Packets() {
		// Extract relevant information from the packet
		port, ip, content := extractPacketInfo(packet)

		// Create a packet string in the expected format for rule matching
		packetString := fmt.Sprintf("Port:%s IP:%s %s", port, ip, content)

		engine.Match(packetString)
	}
}

func extractPacketInfo(packet gopacket.Packet) (string, string, string) {
	// Extract relevant information from the packet (adjust parsing based on your data format)
	// For example, assuming the packet contains TCP and IP layers:
	var port, ip, content string

	if tcpLayer := packet.Layer(gopacket.LayerTypeTCP); tcpLayer != nil {
		tcp := tcpLayer.(*gopacket.Tcp)
		port = fmt.Sprintf("%d", tcp.DstPort)
	}

	if ipLayer := packet.Layer(gopacket.LayerTypeIPv4); ipLayer != nil {
		ip = ipLayer.(*gopacket.IPv4).SrcIP.String()
	}

	if appLayer := packet.ApplicationLayer(); appLayer != nil {
		content = string(appLayer.Payload())
	}

	return port, ip, content
}

func readRulesFromFile(filePath string) ([]ruleengine.Rule, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rules []ruleengine.Rule
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			parts := strings.Fields(line)
			if len(parts) == 3 {
				port, ip, content := parts[0], parts[1], parts[2]
				rules = append(rules, ruleengine.Rule{Port: port, IP: ip, Content: content})
			} else {
				fmt.Println("Invalid rule format:", line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rules, nil
}
