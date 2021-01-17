package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inf = 4294967295
)

// Edge is an edge dog
type Edge struct {
	tail   int
	head   int
	length int
}

// ProbSize is used to define the scope of the all pairs problem
type ProbSize struct {
	numVert  int
	numEdges int
}

// ret -> (smallest smallest path, negative cycle?)
func floyd(edges []Edge, probSize ProbSize) (int, bool) {
	A := buildContainer(probSize)
	vertices := getVertices(edges, probSize.numVert)

	for _, edge := range edges {
		A[edge.tail][edge.head] = edge.length
	}
	for _, vert := range vertices {
		A[vert][vert] = 0
	}

	for k := 0; k < probSize.numVert; k++ {
		for i := 0; i < probSize.numVert; i++ {
			for j := 0; j < probSize.numVert; j++ {
				if A[i][j] > (A[i][k] + A[k][j]) {
					A[i][j] = (A[i][k] + A[k][j])
				}
			}
		}
	}

	// check for neg cycle
	for i := 0; i < probSize.numVert; i++ {
		if A[i][i] < 0 {
			return inf, true
		}
	}

	return findAnswer(A), false
}

func buildContainer(probSize ProbSize) [][]int {
	A := make([][]int, probSize.numVert)
	for i := 0; i < probSize.numVert; i++ {
		A[i] = make([]int, probSize.numVert)
	}
	for i := 0; i < probSize.numVert; i++ {
		for j := 0; j < probSize.numVert; j++ {
			A[i][j] = inf
		}
	}
	return A
}

func getVertices(edges []Edge, size int) []int {
	unique := make(map[int]int, 0)
	res := make([]int, 0)
	for _, edge := range edges {
		tail := edge.tail
		head := edge.head
		if _, exist := unique[tail]; !exist {
			unique[tail] = tail
			res = append(res, tail)
		}
		if _, exist := unique[head]; !exist {
			unique[head] = head
			res = append(res, head)
		}
	}
	return res
}

func findAnswer(A [][]int) int {
	min := inf
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if guess := A[i][j]; guess < min {
				min = guess
			}
		}
	}
	return min
}

func main() {
	fmt.Println(floyd(loadData("./course4/week1/floyd/data1.txt")))
	fmt.Println(floyd(loadData("./course4/week1/floyd/data2.txt")))
	fmt.Println(floyd(loadData("./course4/week1/floyd/data3.txt")))
}

func loadData(filepath string) ([]Edge, ProbSize) {
	items := make([]Edge, 0)
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	row := 1
	probSize := ProbSize{}
	for scanner.Scan() {
		if row == 1 {
			rowSlice := strings.Fields(scanner.Text())
			numVert, err := strconv.Atoi(rowSlice[0])
			check(err)
			numEdges, err := strconv.Atoi(rowSlice[1])
			check(err)
			probSize = ProbSize{
				numVert,
				numEdges,
			}
			row++
			continue
		}
		parseRowIntoEntry(&items, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return items, probSize
}

func parseRowIntoEntry(items *[]Edge, row string) {
	rowSlice := strings.Fields(row)
	tail, err := strconv.Atoi(rowSlice[0])
	check(err)
	head, err := strconv.Atoi(rowSlice[1])
	check(err)
	length, err := strconv.Atoi(rowSlice[2])
	check(err)
	*items = append(*items, Edge{
		tail:   tail - 1,
		head:   head - 1,
		length: length,
	})
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
