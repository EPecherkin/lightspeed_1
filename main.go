package main

import (
	"log"
)

const (
	filename = "ips.txt"
)

func countIPs() {
	hashset := map[byte]map[byte]map[byte]map[byte]bool{}
	totalCount := uint64(0)

	ips := make(chan [4]byte, 1000)
	go readIPs(filename, ips)

	for ip := range ips {
		totalCount += 1

		if hashset[ip[0]] == nil {
			hashset[ip[0]] = map[byte]map[byte]map[byte]bool{}
		}
		m2 := hashset[ip[0]]

		if m2[ip[1]] == nil {
			m2[ip[1]] = map[byte]map[byte]bool{}
		}
		m3 := m2[ip[1]]

		if m3[ip[2]] == nil {
			m3[ip[2]] = map[byte]bool{}
		}
		m4 := m3[ip[2]]

		m4[ip[3]] = true
	}

	uniqCount := uint32(0)
	for _, m2 := range hashset {
		for _, m3 := range m2 {
			for _, m4 := range m3 {
				uniqCount += uint32(len(m4))
			}
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
