package main

import (
	"io"
	"log"
	"os"
)

const (
	readChunkSize = 10 * 1024 * 1024 // 10MB. Even 1 MB is enough, making it bigger doesn't make it faster on my machine, but this reduces the amount of calls, that might help with HDD. Frankly speaking, I would increase it even more for HDD
)

func readIPs(filename string, ips chan [4]byte) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error reading %s: %v\n", filename, err)
		close(ips)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Printf("Error reading %s stats: %v\n", filename, err)
		close(ips)
		return
	}
	totalSize := stat.Size()

	log.Printf("Reading %s %%:\n", filename)

	fileBuffer := make([]byte, readChunkSize)
	processedSize := uint64(0)
	lastReportedPercent := uint8(0)

	ip := [4]byte{}
	ipSegN := 0 // what segment of IP we are building right now
	seg := byte(0)
	segCharN := 0 // what digit of segment should be retrieved next

	for {
		bytesRead, err := file.Read(fileBuffer)
		if err != nil && err != io.EOF {
			log.Printf("Error reading data from %s: %v\n", filename, err)
			close(ips)
			return
		}

		// By-byte parsing, thanks that this big ip_addresses.txt uses ascii encoding
		for i := range bytesRead {
			processedSize++

			char := fileBuffer[i]
			// Any attempt to refactor that just makes it more complicated, since we need to pass multiple values by pointers
			// and still need to handle conditions about malformed file
			if char == '\n' || char == '.' { // NOTE: That won't handle the case when the last \n is missing. But I double checked the file from the task - it is there so we ignore that case
				if ipSegN > 3 {
					log.Printf("%s seems to be malformed at position %d", filename, processedSize)
					close(ips)
					return
				}

				ip[ipSegN] = seg
				seg = 0
				segCharN = 0

				if char == '\n' {
					ipSegN = 0
					ips <- ip
					ip = [4]byte{}
				} else { // '.'
					ipSegN++
				}
			} else {
				num := char - '0' // a little bit of ASII magic
				if num < 0 || num > 9 || segCharN > 2 {
					log.Printf("%s seems to be malformed at position %d", filename, processedSize)
					close(ips)
					return
				}
				seg = seg*10 + num
				segCharN++
			}

			percent := uint8(float64(processedSize) / float64(totalSize) * 100.0)
			if lastReportedPercent != percent {
				log.Printf("  %d%%", percent)
				lastReportedPercent = percent
			}
		}

		if err == io.EOF || bytesRead == 0 {
			close(ips)
			return
		}
	}
}
