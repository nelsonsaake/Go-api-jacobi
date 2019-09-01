package solve

import (
	"math"
)

func product(aj []float64, x []float64) (ajx []float64) {

	if len(aj) != len(x) {
	}

	ajx = make([]float64, len(aj))

	for i, _ := range aj {
		ajx[i] = aj[i] * x[i]
	}

	return
}

func sum(ajx []float64) (result float64) {
	result = 0

	for _, elem := range ajx {
		result += elem
	}

	return
}

func term2(aj []float64, x []float64, i int) (result float64) {

	result = 0 - (sum(product(aj, x)) - aj[i]*x[i])

	return
}

func term1(b []float64, i int) (result float64) {

	result = b[i]

	return
}

func rhs(A [][]float64, x []float64, b []float64, i int) (result float64) {

	result = (1 / A[i][i]) * (term1(b, i) + term2(A[i], x, i))

	return
}

func tolerable(xnew, xold []float64, tol float64) bool {

	for i := range xnew {
		if math.Abs(xnew[i]-xold[i]) > tol {
			return false
		}
	}

	return true
}

func SolveEqns(A [][]float64, x []float64, b []float64, tol float64) ([]float64, [][]float64) {

	tbl := [][]float64{}
	tbl = append(tbl, x)

	for {
		xnew := []float64{}
		for i := range A {
			xnew = append(xnew, rhs(A, x, b, i))
		}

		xold := x
		x = xnew
		tbl = append(tbl, x)

		if tolerable(xold, x, tol) {
			break
		}
	}

	return x, tbl
}
