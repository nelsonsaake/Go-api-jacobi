package solve

import (
	"errors"
	convergence "projects/jacobi/convergence"
	eqns "projects/jacobi/eqns"
	"regexp"
	"strconv"
)

type ResultTable struct {
	Variables []string
	Tbl       [][]float64
}

func toStrs(str string) []string {

	lineExp := ".*\\n"

	lineRegExp := regexp.MustCompile(lineExp)

	return lineRegExp.FindAllString(str, -1)
}

func toEqns(lines []string) (result []eqns.Equation) {

	for _, line := range lines {

		eqn, err := eqns.ToEqn(line)

		if err != nil {
			continue
		}

		result = append(result, eqn)
	}

	return
}

func SolveStr(str string, tolStr string) (res ResultTable, err error) {

	A, b, vars := eqns.ToMat(toEqns(toStrs(str)))

	if !convergence.IsConverge(A) {
		err = errors.New("The system equations will not converge.")
		return
	}

	tol, err := strconv.ParseFloat(tolStr, 64)

	if err != nil {
		err = errors.New("The tolerance should be a real number.")
		return
	}

	x := make([]float64, len(b))

	_, tbl := SolveEqns(A, x, b, tol)

	res = ResultTable{
		Variables: vars,
		Tbl:       tbl,
	}

	return
}
