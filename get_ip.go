package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	readChunkSize = 100 * 1024 * 1024 // 100 MB
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
	processedSize := uint64(0)
	lastReportedPercent := uint8(0)

	var ipBuf bytes.Buffer

	for {
		buffer := make([]byte, readChunkSize)
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Printf("Error reading data from %s: %v\n", filename, err)
			close(ips)
			return
		}

		for i := range bytesRead {
			processedSize += 1
			char := buffer[i]
			if char == '\n' { // TODO: last line break might be missing

				ip := ipBuf.String()
				ipSegments := strings.Split(ip, ".")
				ipBytes := [4]byte{}
				if len(ipSegments) != 4 {
					log.Printf("%s seems to be malformed. Can't parse `%s`", filename, ip)
					close(ips)
					return
				}
				for k := range 4 {
					segByte, err := strconv.ParseUint(ipSegments[k], 10, 8)
					if err != nil {
						log.Printf("%s seems to be malformed. Can't parse segment %s of `%s`: %v", filename, ipSegments[k], ip, err)
						close(ips)
						return
					}
					ipBytes[k] = byte(segByte)
				}

				ips <- ipBytes

				ipBuf = bytes.Buffer{}

				percent := uint8(float64(processedSize) / float64(totalSize) * 100.0)
				if lastReportedPercent != percent {
					log.Printf("  %d%%", percent)
					lastReportedPercent = percent
				}
			} else {
				ipBuf.WriteByte(char)
			}
		}

		if err == io.EOF || bytesRead == 0 {
			close(ips)
			return
		}
	}
}
