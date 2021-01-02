package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Edge is an edge
type Edge struct {
	node1 int
	node2 int
	cost  int
}

// EdgeList is an edge list
type EdgeList []Edge

// UnionFind is the union find
type UnionFind struct {
	root []int
	size []int
}

func (uf *UnionFind) init(size int) *UnionFind {
	uf = new(UnionFind)
	uf.root = make([]int, size)
	uf.size = make([]int, size)

	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) union(p int, q int) {
	qRoot := uf.find(q)
	pRoot := uf.find(p)

	if uf.size[qRoot] < uf.size[pRoot] {
		uf.root[qRoot] = uf.root[pRoot]
		uf.size[pRoot] += uf.size[qRoot]
	} else {
		uf.root[pRoot] = uf.root[qRoot]
		uf.size[qRoot] += uf.size[pRoot]
	}
}

func (uf *UnionFind) find(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}

	for uf.root[p] != p {
		// Path compression
		// Make every other node in path point to its grandparent (thereby halving path length)
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}

	return p
}

func (uf *UnionFind) connected(p int, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UnionFind) numberOfGroups(size int) int {
	leaders := make([]int, 0, size)
	for i := 0; i < size; i++ {
		leaders = append(leaders, uf.find(i))
	}
	return len(unique(leaders))
}

func kClusterings(edges EdgeList, numberOfClustersDesired, numberOfNodes int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})
	uf := new(UnionFind).init(numberOfNodes)

	result := 0
	for _, edge := range edges {
		if !uf.connected(edge.node1, edge.node2) {
			if uf.numberOfGroups(numberOfNodes) == numberOfClustersDesired {
				result = edge.cost
				break
			}
			uf.union(edge.node1, edge.node2)
		}
	}

	return result
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	fmt.Println(kClusterings(loadData("./course3/week2/cluster/data.txt"), 4, 500))
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
			node1: node1 - 1,
			node2: node2 - 1,
			cost:  cost,
		})
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
