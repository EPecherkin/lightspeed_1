package main

import (
	"log"
	"math"
)

const (
	filename = "ips.txt"
)

func countIPs() {
	hashset := map[uint32]bool{}
	totalCount := uint32(0)

	ips := make(chan uint32, 100)
	go readIPs(filename, ips)

	for ip := range ips {
		totalCount += 1
		hashset[ip] = true
	}

	uniqCount := len(hashset)

	log.Printf("Amount of IPs: %d\n", totalCount)
	log.Printf("Amount of uniq IPs: %d\n", uniqCount)
}

func main() {
	monQuit := make(chan bool)
	go trackUsage(monQuit)

	countIPs()

	monQuit <- true
}
