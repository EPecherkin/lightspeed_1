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
  ips := make(chan [4]byte)
	for ip := range ips {
		totalCount += 1
		hash := uint32(0)
		for i := range 4 {
			hash += uint32(ip[i]) * uint32(math.Pow(float64(256), float64(i)))
		}
		hashset[hash] = true
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
