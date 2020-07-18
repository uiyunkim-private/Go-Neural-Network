package matrix

//Matrix object
type Matrix struct {
	shape []int
	dims  int
	data  []float64
}

//SetShape sets shape of a matrix
func (mat *Matrix) SetShape(newShape []int) {
	mat.shape = newShape

}

//GetShape returns shape of a matrix
func (mat *Matrix) GetShape() []int {
	return mat.shape
}

//Zeroes creates empty matrix of given shape
func Zeroes(shape []int) *Matrix {
	mat := &Matrix{shape: shape}

	return mat
}
