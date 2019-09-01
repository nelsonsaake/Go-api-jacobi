package convergence

import "math"

func sum(row []float64) (result float64) {
	result = 0

	for _, elem := range row {
		result += elem
	}

	return
}

func willConverge(row []float64, i int) bool {

	return math.Abs(sum(row)-row[i]) <= math.Abs(row[i])
}

func IsConverge(mat [][]float64) bool {

	for i, row := range mat {
		if !willConverge(row, i) {
			return false
		}
	}

	return true
}
