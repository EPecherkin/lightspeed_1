package main

import (
	"fmt"
	"log"
)

type Node struct {
	digits   []byte
	children []*Node
}

func add(root *Node, ipDigits *[10]byte) byte {
	node := root
	digits := ipDigits[:]

	inserted := byte(0)

	for len(digits) > 0 {
		nextNode := (*Node)(nil)
		for i := range len(node.children) {
			tnode := node.children[i]
			digitsLen := len(tnode.digits)

			// Find first difference position
			k := 0
			for k < digitsLen {
				if digits[k] != tnode.digits[k] {
					break
				}
				k++
			}

			if k == 0 { // no match at all, check other nodes
				continue
			}

			if (k == digitsLen) && (len(digits) == len(tnode.digits)) { // exact match
				return 0
			} else { // target node found
				digits = digits[k:]
				nextNode = tnode
				if k < digitsLen {
					tdigits := tnode.digits[:k]
					cdigits := tnode.digits[k:]
					cnode := Node{digits: cdigits, children: tnode.children}
					tnode.digits = tdigits
					tnode.children = []*Node{&cnode}
				}
				break
			}
		}
		if nextNode == nil {
			inserted = 1
			cnode := Node{digits: digits}
			nextNode = &cnode
			node.children = append(node.children, nextNode)
			break
		}
		node = nextNode
	}
	return inserted
}

func printTree(root *Node, count uint64) {
	line := [][]*Node{[]*Node{root}}

	for _ = range count {
		nextline := [][]*Node{}
		for i := range line {
			fmt.Print("[")
			for k := range line[i] {
				fmt.Printf("%v", line[i][k].digits)
				nextline = append(nextline, line[i][k].children)
			}
			fmt.Print("] ")
		}
		line = nextline
		fmt.Println("")
	}
}

func buildTree(ips chan *[10]byte) {
	root := Node{}
	totalCount := uint64(0)
	uniqCount := uint64(0)

	for ip := range ips {
		totalCount++
		uniqCount += uint64(add(&root, ip))
	}
	printTree(&root, uniqCount)

	log.Printf("Amount of IPs: %d\n", totalCount)
	log.Printf("Amount of uniq IPs: %d\n", uniqCount)
}
