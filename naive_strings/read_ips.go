package main

import (
	"bytes"
	"io"
	"iter"
	"os"
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

		var ip bytes.Buffer
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
				if char == '\n' { // TODO: last empty line might be missing
					percent := uint8(float64(processedSize) / float64(totalSize) * 100.0)
					if !yield(ip.String(), IPReadingProgress{percent: percent}) {
						return
					}
					ip = bytes.Buffer{}
				} else {
					ip.WriteByte(char)
				}
			}

			if err == io.EOF || bytesRead == 0 {
				return
			}
		}
	}
}
