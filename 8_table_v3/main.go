package main

import (
	"log"
	"os"
)

const (
	ipsChannelSize = 10000
)

func countIPs() {
	filename := os.Getenv("IPS_FILE")
	if len(filename) == 0 {
		log.Println("You should specify location of the file with IPs with IPS_FILE environment variable")
		return
	}

	ips := make(chan [4]byte, 10000)
	go readIPs(filename, ips)

	// We represent the last 256 digits as 32*8 bits, going to use binary operations to check/set if number is present
	table := [256][256][256][32]byte{}

	totalCount := uint64(0)
	uniqCount := uint64(0)

	for ip := range ips {
		totalCount++

		// segment % 32 to get the proper basket
		seg4Basket := byte(ip[3] % 32)
		// 1 << (segment / 32) to get the bit that represents the number in the basket
		seg4Bit := byte(1 << (ip[3] / 32))

		if (table[ip[0]][ip[1]][ip[2]][seg4Basket] & seg4Bit) == 0 {
			uniqCount++
			table[ip[0]][ip[1]][ip[2]][seg4Basket] |= seg4Bit
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
