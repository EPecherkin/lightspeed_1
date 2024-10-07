package main

import (
  "fmt"
	"io"
	"iter"
	"os"
  "math"
)

const (
	readChunkSize = 100 * 1024 * 1024 // 100 MB
)

type IPReadingProgress struct {
	percent uint8
	err     error
}

func getIP(filename string) iter.Seq2[[4]byte, IPReadingProgress] {
	return func(yield func([4]byte, IPReadingProgress) bool) {
    emptyVal := [4]byte{}
		file, err := os.Open(filename)
		if err != nil {
			yield(emptyVal, IPReadingProgress{err: err})
			return
		}
		defer file.Close()
		stat, err := file.Stat()
		if err != nil {
			yield(emptyVal, IPReadingProgress{err: err})
			return
		}
		totalSize := stat.Size()
		processedSize := uint64(0)

    ip := ""
		for {
			buffer := make([]byte, readChunkSize)
			bytesRead, err := file.Read(buffer)
			if err != nil && err != io.EOF {
				yield(emptyVal, IPReadingProgress{err: err})
				return
			}

			for i := range bytesRead {
				processedSize += 1
				char := buffer[i]
				if char == '\n' || err == io.EOF || bytesRead == 0
          segIndex = 0
					percent := uint8(float64(processedSize) / float64(totalSize) * 100.0)
					if !yield(ip, IPReadingProgress{percent: percent}) {
						return
					}
					ip = [4]byte{}
				} else if char == '.' {
          var segVal byte
          for k := range len(seg) {
            segVal += seg[k]*byte(math.Pow(10.0, float64(len(seg)-k-1)))
          }
					ip[segIndex] = segVal
          segIndex += 1
					seg = []byte{}
				} else {
					seg = append(seg, char - '0')
				}
			}

			if bytesRead == 0 || err == io.EOF {
				break
			}
		}
	}
}
