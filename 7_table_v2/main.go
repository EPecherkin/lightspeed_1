package main

import (
	"log"
	"os"
)

func countIPs() {
	ips := make(chan [4]byte, 10000)
	filename := os.Getenv("IPS_FILE")
	go readIPs(filename, ips)

	table := [256][256][256][32]byte{}

	totalCount := uint64(0)
	uniqCount := uint64(0)
	averageChannelSize := float64(0)
	maxChannelSize := uint16(0)

	for ip := range ips {
		averageChannelSize = float64((uint64(averageChannelSize)*totalCount + uint64(len(ips)))) / float64(totalCount+1)
		if uint16(len(ips)) > maxChannelSize {
			maxChannelSize = uint16(len(ips))
		}
		totalCount++
		if (table[ip[0]][ip[1]][ip[2]][ip[3]%32] & (1 << (ip[3] / 32))) == 0 {
			uniqCount++
			table[ip[0]][ip[1]][ip[2]][ip[3]%32] |= (1 << (ip[3] / 32))
		}
	}

	log.Printf("Amount of IPs: %d\n", totalCount)
	log.Printf("Amount of uniq IPs: %d\n", uniqCount)
	log.Printf("Average channel size: %0.2f\n", averageChannelSize)
	log.Printf("Max channel size: %d\n", maxChannelSize)
}

func main() {
	monQuit := make(chan bool)
	go trackUsage(monQuit)

	countIPs()

	monQuit <- true
}
