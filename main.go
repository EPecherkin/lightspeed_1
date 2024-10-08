package main

import (
)

const (
	filename = "ips_s.txt"
)

func countIPs() {
	ips := make(chan uint32, 1000)
	go readIPs(filename, ips)

  buildTree(ips)
}

func main() {
	monQuit := make(chan bool)
	go trackUsage(monQuit)

	countIPs()

	monQuit <- true
}
