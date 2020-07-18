package main

import (
	"fmt"
	"godeep/matrix"
)

func main() {
	shape := []int{111, 111}

	var mat *matrix.Matrix
	mat = matrix.Zeroes(shape)

	fmt.Println(mat)
	newShape := []int{120, 120}
	mat.SetShape(newShape)
	fmt.Println(mat)

}
