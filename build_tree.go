package main

import (
	// "log"
	// "sync"
)

type Node struct {
	children map[byte]*Node
	// mu       sync.Mutex
}

const (
)

func add(root *Node, ip [4]byte) {
	node := root

	for i := range 4 {
		// mu := &node.mu
		// mu.Lock()
		if node.children[ip[i]] == nil {
			tnode := Node{children: map[byte]*Node{}}
			node.children[ip[i]] = &tnode
		}
		node = node.children[ip[i]]
		// mu.Unlock()
	}
}

// func buildTreeRoutine(ips chan [4]byte, wg *sync.WaitGroup) {
func buildTreeRoutine(ips chan [4]byte) {
  // wg.Add(1)
  // defer wg.Done()

	root := Node{children: map[byte]*Node{}}

	for ip := range ips {
    add(&root, ip)
	}
}

func buildTree(ips chan [4]byte) {
	// wg := sync.WaitGroup{}

  // for _ = range concurrency {
  //   go buildTreeRoutine(ips, &wg)
  // }
  // for _ = range concurrency {
    buildTreeRoutine(ips)
  // }

	// wg.Wait()

	// log.Printf("Amount of IPs: %d\n", 0)
	// log.Printf("Amount of uniq IPs: %d\n", uniqCount)
}
