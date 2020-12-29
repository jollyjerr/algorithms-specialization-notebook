package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Edge is an edge
type Edge struct {
	node1 Node
	node2 Node
	cost  int
}

// EdgeList is an edge list
type EdgeList []Edge

// Node is a node
type Node struct {
	vertex int
	leader int
}

// Group is a group
type Group struct {
	size  int
	nodes []*Node
}

// UnionFind supports find and union opperations
type UnionFind struct {
	allNodes []Node
	groups   map[int]Group
}

func (u UnionFind) init(edges *EdgeList) {
	for _, edge := range *edges {
		if _, exist := u.groups[edge.node1.vertex]; !exist {
			u.allNodes[edge.node1.vertex] = edge.node1
			u.groups[edge.node1.vertex] = Group{
				size:  1,
				nodes: []*Node{&u.allNodes[edge.node1.vertex]},
			}
		}
		if _, exist := u.groups[edge.node2.vertex]; !exist {
			u.allNodes[edge.node2.vertex] = edge.node2
			u.groups[edge.node2.vertex] = Group{
				size:  1,
				nodes: []*Node{&u.allNodes[edge.node2.vertex]},
			}
		}
	}
}

func (u *UnionFind) find(vertex int) int {
	return u.allNodes[vertex].leader
}

func (u UnionFind) union(leader1, leader2 int) {
	if group1, group2 := u.groups[leader1], u.groups[leader2]; group1.size > group2.size {
		for _, node := range group2.nodes {
			node.leader = leader1
		}
		group1.nodes = append(group1.nodes, group2.nodes...)
		group1.size = len(group1.nodes)
		u.groups[leader1] = group1
		delete(u.groups, leader2)
	} else {
		for _, node := range group1.nodes {
			node.leader = leader2
		}
		group2.nodes = append(group2.nodes, group1.nodes...)
		group2.size = len(group2.nodes)
		u.groups[leader2] = group2
		delete(u.groups, leader1)
	}
}

func kClusterings(edges EdgeList, numberOfClustersDesired int) int {
	uf := UnionFind{
		allNodes: make([]Node, len(edges)),
		groups:   make(map[int]Group),
	}
	uf.init(&edges)
	uf.union(499, 500)
	uf.union(498, 500)
	// fmt.Println(uf)
	fmt.Println(uf.find(498))
	fmt.Println(uf.find(499))
	fmt.Println(uf.find(402))
	return 2
}

func main() {
	fmt.Println(kClusterings(loadData("./course3/week2/cluster/data.txt"), 4))
}

func loadData(filepath string) EdgeList {
	data := make(EdgeList, 0, 500)
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parseRowIntoEntry(&data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func parseRowIntoEntry(container *EdgeList, row string) {
	rowSlice := strings.Fields(row)
	if len(rowSlice) > 1 {
		node1, err := strconv.Atoi(rowSlice[0])
		check(err)
		node2, err := strconv.Atoi(rowSlice[1])
		check(err)
		cost, err := strconv.Atoi(rowSlice[2])
		check(err)
		*container = append(*container, Edge{
			node1: Node{vertex: node1, leader: node1},
			node2: Node{vertex: node2, leader: node2},
			cost:  cost,
		})
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
