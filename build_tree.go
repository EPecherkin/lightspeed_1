package main

import (
  "log"
  // "sync"
)

type Node struct {
  seg byte
  children []*Node
}

const (
  maxPos = 32 // 32 bits max
)

func add(root *Node, ip [4]byte) uint64 {
  node := root
  isNew := uint64(0)
  // log.Printf("%v ", ip)

  for i := range 4 {
    // log.Printf("%v %v, %v", ip[i], i, len(node.children))
    k := 0
    found := false
    for k < len(node.children) {
      tnode := node.children[k]
      if tnode.seg == ip[i] {
        // log.Printf("found %v at %v", tnode.seg, k)
        node = tnode
        found = true
        break
      }
      k += 1
    }
    if !found {
      tnode := Node{seg: ip[i]}
      // log.Printf("created %v for %v at %v", tnode.seg, node.seg, k)
      node.children = append(node.children, &tnode)
      node = &tnode
      isNew = 1
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
