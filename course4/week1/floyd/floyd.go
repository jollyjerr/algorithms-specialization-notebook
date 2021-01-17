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
	inf = int(^uint(0) >> 1)
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

func floyd(edges []Edge, probSize ProbSize) int {
	A := buildContainer(probSize) // 3D array (indexed by i,j,k)

	// Pre-process for base cases
	for i := 0; i < probSize.numVert; i++ {
		for j := 0; j < probSize.numVert; j++ {
			if i == j {
				A[i][j][0] = 0
			} else {
				for _, edge := range edges {
					if edge.tail == i && edge.head == j { //|| (edge.tail == j && edge.head == i) ????
						fmt.Println(edge.length)
						A[i][j][0] = edge.length
					} else {
						A[i][j][0] = inf
					}
				}
			}
		}
	}

	// Fill the remaining table
	for k := 1; k < probSize.numVert; k++ {
		for i := 1; i < probSize.numVert; i++ {
			for j := 1; j < probSize.numVert; j++ {
				A[i][j][k] = minOfStep(A, i, j, k)
			}
		}
	}

	fmt.Println(A)

	return 2
}

func buildContainer(probSize ProbSize) [][][]int {
	A := make([][][]int, probSize.numVert)
	for i := 0; i < probSize.numVert; i++ {
		A[i] = make([][]int, probSize.numVert)
		for j := 0; j < probSize.numVert; j++ {
			A[i][j] = make([]int, probSize.numVert)
		}
	}
	return A
}

func minOfStep(A [][][]int, i, j, k int) int {
	inherit := A[i][j][k-1]
	adopt := A[i][k][k-1] + A[k][j][k-1]
	if inherit > adopt {
		return adopt
	}
	return inherit
}

func main() {
	fmt.Println(floyd(loadData("./course4/week1/floyd/smallData.txt")))
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
