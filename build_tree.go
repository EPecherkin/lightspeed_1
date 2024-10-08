package main

import (
  "fmt"
  "errors"
  "log"
  "strconv"
  // "sync"
)

type Node struct {
  data *string
  children []*Node
}

const (
  maxPos = 32 // 32 bits max
)

func add(root *Node, initData *string) error {
  node := root
  data := initData

  for {
    targetNode := (*Node{})(nil)
    for i := range len(node.children) {
      tnode := node.children[i]
      datalen := len(data)
      if len(tnode.data) < datalen {
        datalen = len(tnode.data)
      }

      // Find first difference position
      k := 0
      for k < datalen {
        if data[k] != tnode.data[k] {
          break
        }
        k += 1
      }

      if (k == datalen) && (len(data) == len(tnode.data)) { // exact match
        return
      } else if k == datalen { // one is substring
        if len(data) < len(tnode.data) { // data is subset of tnode.data
          tnode = 
        } else { // tnode.data a subset of what we are looking for

        }
      }

      // if k == 0 { // no match at all, check other nodes
    }
  }

  if len(node.children) == 0 {
    newNode = Node{data}
    append(node.children, &newNode)
  } else {
    tn := (*Node{})(nil)
    for i := range len(node.children) {
      bit = 1 << node.from
      if (node.children[i].data & bit) == (data & bit) {
        tn := *node.children[i]
        break
      }
    }
    if tn != nil {

    } else {
      newNode = Node{data, node.to, maxPos}
      append(node.children, &newNode)
    }
  }
  return nil
}

func buildTree(ips chan uint32) uint64 {
  root := Node{}

	for ip := range ips {
    ipStr := strconv.FormatInt(ip, 2)
    err := add(&root, &ipStr)
    if err != nil {
      log.Printf("Can't add ip %s to Tree %v\n%s\n", pip, root, err)
    }
  }
	//
	// log.Printf("Amount of IPs: %d\n", totalCount)
	// log.Printf("Amount of uniq IPs: %d\n", uniqCount)
  return 1
}
