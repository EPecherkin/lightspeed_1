package main

import (
	"io"
	"log"
	"os"
)

const (
	readChunkSize = 100 * 1024 * 1024 // 100 MB
)

func readIPs(filename string, ips chan uint32) {
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
	processedSize := uint64(0)
	lastReportedPercent := uint8(0)

	fileBuffer := make([]byte, readChunkSize)
	ip := uint32(0)
	ipSegN := 0
	seg := byte(0)
	segCharN := 0

	for {
		bytesRead, err := file.Read(fileBuffer)
		if err != nil && err != io.EOF {
			log.Printf("Error reading data from %s: %v\n", filename, err)
			close(ips)
			return
		}

		for i := range bytesRead {
			processedSize += 1

			char := fileBuffer[i]
			if char == '\n' || char == '.' { // TODO: last line break might be missing
				if ipSegN > 3 {
					log.Printf("%s seems to be malformed at position %d", filename, processedSize)
					close(ips)
					return
				}

				ip = ip << 8 + uint32(seg)
				seg = 0
				segCharN = 0

				if char == '\n' {
					ipSegN = 0
					ips <- ip
					ip = uint32(0)
				} else { // '.'
					ipSegN += 1
				}
			} else {
				num := char - '0'
				if num < 0 || num > 9 || segCharN > 2 {
					log.Printf("%s seems to be malformed at position %d", filename, processedSize)
					close(ips)
					return
				}
				seg = seg*10 + num
				segCharN += 1
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
