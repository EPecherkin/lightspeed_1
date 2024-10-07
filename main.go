package main

import (
	"fmt"
	"log"
	"math"
)

const (
	filename = "ips_s.txt"
)

func main() {
	fmt.Printf("Processing: 0%%")

	hashset := map[uint32]bool{}
	totalCount := uint32(0)
	for ip, progress := range getIP(filename) {
		if progress.err != nil {
			log.Fatalf("Error reading file: %v", progress.err)
			return
		}
		totalCount += 1
		hash := uint32(0)
		for i := range 4 {
			hash += uint32(ip[i]) * uint32(math.Pow(float64(256), float64(i)))
		}
		hashset[hash] = true
		fmt.Printf("\rProcessing: %d%%", progress.percent)
	}
	fmt.Printf("\n")

	uniqCount := len(hashset)

	log.Printf("Amount of IPs: %d\n", totalCount)
	log.Printf("Amount of uniq IPs: %d\n", uniqCount)
}
