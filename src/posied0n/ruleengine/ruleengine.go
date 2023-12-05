package ruleengine

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// Rule represents a Snort-like rule with port, IP, and content conditions
type Rule struct {
	Port    string
	IP      string
	Content string
}

// RuleEngine represents the Snort-like rule engine
type RuleEngine struct {
	Rules []Rule
}

// NewRuleEngine creates a new RuleEngine with the provided rules
func NewRuleEngine(rules []Rule) *RuleEngine {
	return &RuleEngine{Rules: rules}
}

// Match checks if the given packet matches any rule in the rule engine
func (re *RuleEngine) Match(packet string) {
	for _, rule := range re.Rules {
		if re.matchPort(packet, rule.Port) && re.matchIP(packet, rule.IP) && re.matchContent(packet, rule.Content) {
			fmt.Println("Packet matched rule:", rule)
			// Perform action or logging here
		}
	}
}

// matchPort checks if the packet matches the specified port condition in the rule
func (re *RuleEngine) matchPort(packet string, port string) bool {
	// Extract port from the packet (adjust parsing based on your data format)
	// For example, assuming the packet contains "Port:1234", you can use a regex:
	reStr := fmt.Sprintf("Port:(%s)", port)
	rePattern := regexp.MustCompile(reStr)
	return rePattern.MatchString(packet)
}

// matchIP checks if the packet matches the specified IP condition in the rule
func (re *RuleEngine) matchIP(packet string, ip string) bool {
	// Extract IP from the packet (adjust parsing based on your data format)
	// For example, assuming the packet contains "IP:192.168.1.1", you can use a regex:
	reStr := fmt.Sprintf("IP:(%s)", ip)
	rePattern := regexp.MustCompile(reStr)
	return rePattern.MatchString(packet)
}

// matchContent checks if the packet matches the specified content condition in the rule
func (re *RuleEngine) matchContent(packet string, content string) bool {
	// Convert packet hex string to ASCII string
	hexBytes, err := hex.DecodeString(packet)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return false
	}
	asciiPacket := string(hexBytes)

	// Check if the content is present in the packet
	return strings.Contains(asciiPacket, content)
}
