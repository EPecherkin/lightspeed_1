package main

import (
	"log"
	// "sync"
)

type Node struct {
	children map[byte]*Node
}

const (
	maxPos = 32 // 32 bits max
)

func add(root *Node, ip [4]byte) uint64 {
	node := root
	isNew := uint64(0)

	for i := range 4 {
		if node.children[ip[i]] == nil {
			isNew = 1
			tnode := Node{children: map[byte]*Node{}}
			tnode.children[ip[i]] = &tnode
			node = &tnode
		} else {
			node = node.children[ip[i]]
		}
	}
	return isNew
}

func buildTree(ips chan [4]byte) {
	root := Node{}

	totalCount := uint64(0)
	uniqCount := uint64(0)

	for ip := range ips {
		totalCount += 1
		uniqCount += add(&root, ip)
	}

	log.Printf("Amount of IPs: %d\n", totalCount)
	log.Printf("Amount of uniq IPs: %d\n", uniqCount)
}
