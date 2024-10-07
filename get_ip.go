package main

import (
	"io"
	"iter"
	"os"
	"strconv"
)

const (
	readChunkSize = 100 * 1024 * 1024 // 100 MB
)

type IPReadingProgress struct {
	percent uint8
	err     error
}

func getIP(filename string) iter.Seq2[string, IPReadingProgress] {
	return func(yield func(string, IPReadingProgress) bool) {

		file, err := os.Open(filename)
		if err != nil {
			yield("", IPReadingProgress{err: err})
			return
		}
		defer file.Close()
		stat, err := file.Stat()
		if err != nil {
			yield("", IPReadingProgress{err: err})
			return
		}
		totalSize := stat.Size()
		processedSize := uint64(0)

		var ip []byte
		var seg []byte
		for {
			buffer := make([]byte, readChunkSize)
			bytesRead, err := file.Read(buffer)
			if err != nil && err != io.EOF {
				yield("", IPReadingProgress{err: err})
				return
			}

			for i := range bytesRead {
				processedSize += 1
				char := buffer[i]
				// TODO: there might not be another line
				if char == '\n' {
					percent := uint8(float64(processedSize) / float64(totalSize) * 100.0)
					if !yield(string(ip), IPReadingProgress{percent: percent}) {
						return
					}
					ip = []byte{}
				} else if char == '.' {
					i, err := strconv.Atoi(string(seg))
					if err != nil {
						yield("", IPReadingProgress{err: err})
						return
					}
					ip = append(ip, byte(i))
					seg = []byte{}
				} else {
					seg = append(seg, char)
				}
			}

			if bytesRead == 0 || err == io.EOF {
				break
			}
		}
	}
}
