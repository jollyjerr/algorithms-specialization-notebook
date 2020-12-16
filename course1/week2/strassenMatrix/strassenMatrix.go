package main

func strassen(left [][]int, right [][]int) [][]int {
	// Base case
	if len(left) <= 1 {
		return [][]int{
			{left[0][0] * right[0][0]},
		}
	}

	// TODO this still has some bugs

	// a, b, c, d := split(left)
	// e, f, g, h := split(right)

	// p1 := strassen(a, sub(f, h))
	// p2 := strassen(add(a, b), h)
	// p3 := strassen(add(c, d), e)
	// p4 := strassen(d, sub(g, e))
	// p5 := strassen(add(a, d), add(e, h))
	// p6 := strassen(sub(b, d), add(g, h))
	// p7 := strassen(sub(a, c), add(e, f))

	// fmt.Println(p1)
	// fmt.Println(p2)
	// fmt.Println(p3)
	// fmt.Println(p4)
	// fmt.Println(p5)
	// fmt.Println(p6)
	// fmt.Println(p7)

	// half := len(left) / 2
	// result := makeEmpty2DSlice(len(left))

	// for i := 0; i < half; i++ {
	// 	for j := 0; j < half; j++ {
	// 		fmt.Println(result)
	// 	}
	// }

	return [][]int{{2}}
}

func split(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {
	// We are only working with even matrices
	height := len(matrix)
	halfHeight := height / 2

	var (
		topLeft     = makeEmpty2DSlice(halfHeight)
		topRight    = makeEmpty2DSlice(halfHeight)
		bottomLeft  = makeEmpty2DSlice(halfHeight)
		bottomRight = makeEmpty2DSlice(halfHeight)
	)

	for i := 0; i < halfHeight; i++ {
		for j := 0; j < halfHeight; j++ {
			topLeft[i][j] = matrix[i][j]
			topRight[i][j] = matrix[i][j+halfHeight]
			bottomLeft[i][j] = matrix[i+halfHeight][j]
			bottomRight[i][j] = matrix[i+halfHeight][j+halfHeight]
		}
	}

	return topLeft, topRight, bottomLeft, bottomRight
}

func add(left [][]int, right [][]int) [][]int {
	half := len(left) / 2
	res := makeEmpty2DSlice(half)

	for i := 0; i < half; i++ {
		for j := 0; j < half; j++ {
			res[i][j] = (left[i][j] + right[i][j])
		}
	}

	return res
}

func sub(left [][]int, right [][]int) [][]int {
	half := len(left) / 2
	res := makeEmpty2DSlice(half)

	for i := 0; i < half; i++ {
		for j := 0; j < half; j++ {
			res[i][j] = (left[i][j] - right[i][j])
		}
	}

	return res
}

// IDK if this is the best way to do this but Go was being silly
func makeEmpty2DSlice(len int) [][]int {
	a := make([][]int, len)
	for i := range a {
		a[i] = make([]int, len)
	}
	return a
}
