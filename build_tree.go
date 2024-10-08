package main

import (
	"log"
	"sync"
)

type Node struct {
	children map[byte]*Node
	mu       sync.Mutex
}

const (
	maxPos = 32 // 32 bits max
)

func add(root *Node, ip [4]byte, wg sync.WaitGroup) {
	node := root
	wg.Add(1)

	for i := range 4 {
		mu := node.mu
		mu.Lock()
		if node.children[ip[i]] == nil {
			tnode := Node{children: map[byte]*Node{}}
			tnode.children[ip[i]] = &tnode
			node = &tnode
		} else {
			node = node.children[ip[i]]
		}
		mu.Unlock()
	}
	wg.Done()
}

func buildTree(ips chan [4]byte) {
	wg := sync.WaitGroup{}

	root := Node{}

	totalCount := uint64(0)
	uniqCount := uint64(0)

	for ip := range ips {
		totalCount += 1
		go add(&root, ip, wg)
	}
	wg.Wait()

	log.Printf("Amount of IPs: %d\n", totalCount)
	log.Printf("Amount of uniq IPs: %d\n", uniqCount)
}
