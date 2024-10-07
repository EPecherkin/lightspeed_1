package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
)

const (
	// totalFileSize = 1 * 1024 * 1024 // 1 MB for tests
	totalFileSize = 1 * 1024 * 1024 * 1024 // 1 GB is enough
	batchSize     = 100000                 // Amount to store in memory before flushing
	outputFile    = "ips.txt"
	duplicate     = 1000 // 1 / 1000 chance to duplicate an IP
)

func generateRandomIP() string {
	ip := make(net.IP, 4)
	rand.Read(ip)
	return ip.String()
}

func main() {
	file, err := os.Create(outputFile)
	if err != nil {
		log.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	totalBytesWritten := uint64(0)

	for totalBytesWritten < totalFileSize {
		// Generate IPs and add them to the batch
		ipBatch := make([]string, batchSize)
		for i := range batchSize {
			if i > 1 && rand.Intn(duplicate) == 1 { // ensure diplicates
				ipBatch[i] = ipBatch[i-1]
			} else {
				ipBatch[i] = generateRandomIP()
			}
		}

		batchData := strings.Join(ipBatch[:], "\n")

		// Write the batch to the file
		bytesWritten, err := file.WriteString(batchData)
		if err != nil {
			log.Printf("Error writing to file: %v\n", err)
			return
		}

		totalBytesWritten += uint64(bytesWritten)
		fmt.Printf("\rBytes written: %d / %d", totalBytesWritten, totalFileSize)
	}
	fmt.Printf("\n")

	log.Println("Finished generating ips.txt random IP addresses.")
}
