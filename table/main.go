package main

import (
	"log"
	"os"
)

func countIPs() {
	ips := make(chan [4]byte, 1000)
	filename := os.Getenv("IPS_FILE")
	go readIPs(filename, ips)

	table := [256][256][256][256]byte{}

	totalCount := uint64(0)
	uniqCount := uint64(0)

	for ip := range ips {
		totalCount++
		if table[ip[0]][ip[1]][ip[2]][ip[3]] == 0 {
			uniqCount++
			table[ip[0]][ip[1]][ip[2]][ip[3]] = 1
		}
	}

	log.Printf("Amount of IPs: %d\n", totalCount)
	log.Printf("Amount of uniq IPs: %d\n", uniqCount)
}

func main() {
	monQuit := make(chan bool)
	go trackUsage(monQuit)

	countIPs()

	monQuit <- true
}
