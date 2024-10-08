package main

import ()

const (
	filename = "ips_b.txt"
)

func countIPs() {
	ips := make(chan [4]byte, 1000)
	go readIPs(filename, ips)

	buildTree(ips)
}

func main() {
	monQuit := make(chan bool)
	go trackUsage(monQuit)

	countIPs()

	monQuit <- true
}
