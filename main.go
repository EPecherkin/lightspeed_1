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
  // fmt.Printf("Processing: 0%%")

  var hashset map[uint32]bool
  for ip, progress := range getIP(filename) {
    fmt.Println(ip)
    if progress.err != nil {
      log.Fatalf("Error reading file: %v", progress.err)
      return
    }
    var hash uint32
    for i := range 4 {
      hash += uint32(ip[i]) * uint32(math.Pow(float64(256), float64(i)))
    }
    hashset[hash] = true
    // fmt.Printf("\rProcessing: %d%%", progress.percent)
  }
  fmt.Printf("\n")

  count := len(hashset)

  log.Printf("Amount of uniq IPs: %d\n", count)
}

