package eqns

import (
	"sort"
)

//
func contains(list []string, val string) bool {

	for _, str := range list {
		if str == val {
			return true
		}
	}

	return false
}

func countNumberOfVariables(eqns []Equation) (int, []string) {

	variables := []string{}

	for _, eqn := range eqns {
		for _, term := range eqn.variableTerms {
			if !contains(variables, term.variable) {
				variables = append(variables, term.variable)
			}
		}
	}

	return len(variables), variables
}

func isDeterministic(eqns []Equation) bool {

	count, _ := countNumberOfVariables(eqns)

	return count == len(eqns)
}

//
func sortEqn(eqn Equation) Equation {

	// arrange variable according to particular scheme

	sort.SliceStable(eqn.variableTerms, func(i, j int) bool {
		return eqn.variableTerms[i].variable < eqn.variableTerms[j].variable
	})

	return eqn
}

//
func toRow(eqn Equation) (row []float64) {

	for _, term := range eqn.variableTerms {
		row = append(row, term.coefficient)
	}

	return
}

func coefficientOfTerm(eqn Equation, variable string) (coefficient float64) {

	// finds all variable terms with the variable part same a the one specified
	// sums all their coefficients and return it

	for _, term := range eqn.variableTerms {
		if term.variable == variable {
			coefficient = coefficient + term.coefficient
		}
	}

	return
}

func toRowOrderByShceme(eqn Equation, variables []string) (row []float64) {

	for _, variable := range variables {
		row = append(row, coefficientOfTerm(eqn, variable))
	}

	return
}

func ToMat(eqns []Equation) (A [][]float64, b []float64, variables []string) {

	_, variables = countNumberOfVariables(eqns)

	for _, eqn := range eqns {
		A = append(A, toRowOrderByShceme(eqn,variables))
		b = append(b, eqn.constant)
	}

	return
}
