package main

import (
	"log"
)

const (
	filename = "ips.txt"
)

func countIPs() {
	log.Println("Processing %:")

	hashset := map[string]bool{}
	totalCount := uint64(0)
	lastReportedPercent := uint8(0)
	for ip, progress := range getIP(filename) {
		if progress.err != nil {
			log.Fatalf("Error reading file: %v", progress.err)
			return
		}
		totalCount += 1
		hashset[ip] = true
		if progress.percent != lastReportedPercent {
			log.Printf("  %d%%\n", progress.percent)
			lastReportedPercent = progress.percent
		}
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
