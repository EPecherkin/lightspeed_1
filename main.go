package main

import ()

const (
	filename = "ips.txt"
  concurrency = 10
)

func countIPs() {
	ips := make(chan [4]byte, 10)
	go readIPs(filename, ips)

	buildTree(ips)
}

func main() {
	monQuit := make(chan bool)
	go trackUsage(monQuit)

	countIPs()

	monQuit <- true
}
