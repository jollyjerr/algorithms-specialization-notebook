package main

import (
	"fmt"
)

// Node is a node
type Node struct {
	vertex   int
	edgesTo  []int
	explored bool
}

func tsat(graph []Node) bool {
	fmt.Println(graph)
	return false
}

func main() {
	fmt.Println(tsat(buildGraph("course4/week4/2sat/smallData.txt")))
}
