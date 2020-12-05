package main

import "fmt"

func strassen(left [][]int, right [][]int) [][]int {
	// Base case
	if len(left) == 1 {
		return [][]int{
			{left[0][0] * right[0][0]},
		}
	}

	a, b, c, d := split(left)
	e, f, g, h := split(right)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	p1 := strassen(a, sub(f, h))
	p2 := strassen(add(a, b), h)
	p3 := strassen(add(c, d), e)
	p4 := strassen(d, sub(g, e))
	p5 := strassen(add(a, d), add(e, h))
	p6 := strassen(sub(b, d), add(g, h))
	p7 := strassen(sub(a, c), add(e, f))

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println(p5)
	fmt.Println(p6)
	fmt.Println(p7)

	return [][]int{{2}}
}

func split(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {
	rowLen := len(matrix)

	half := rowLen / 2

	A := matrix[:half]
	B := matrix[half:]

	return [][]int{A[0][:half]}, [][]int{A[0][half:]}, [][]int{B[0][:half]}, [][]int{B[0][half:]}
}

func add(left [][]int, right [][]int) [][]int {
	// res := make([][]int, len(left))

	for i := 0; i < len(left); i++ {
		for j := 0; j < len(left); j++ {
			// res[i][j] = (left[i][j] + right[i][j])
			// fmt.Println(res)
		}
	}
	return [][]int{{2}}
}

func sub(left [][]int, right [][]int) [][]int {
	// res := make([][]int, len(left))
	for i := 0; i < len(left); i++ {
		for j := 0; j < len(left); j++ {
			// res[i][j] = (left[i][j] - right[i][j])
			// fmt.Println(res)
		}
	}
	return [][]int{{2}}
}
