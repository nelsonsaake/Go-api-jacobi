package eqns

import (
	"errors"
	"regexp"
	"strconv"
)

type Term struct {
	coefficient float64
	variable    string
}

func isValidTerm(termText string) (err error) {

	ismatch, err := regexp.MatchString(termExp, termText)
	if err != nil {
		return err
	}

	if !ismatch {
		err = errors.New(
			"Text is not a term." +
				"\nTerm should take the form +2a or 3c or d.")

		return err
	}

	return err
}

func numberPart(termText string) (num float64, err error) {

	coefficientRegExp := regexp.MustCompile(numberExp)

	numStr := coefficientRegExp.FindString(termText)

	if numStr == "" {
		num = 0
		return
	}

	num, err = strconv.ParseFloat(numStr, 64)

	return
}

func signPart(termText string) (sign string) {

	signRegExp := regexp.MustCompile(signExp)

	sign = signRegExp.FindString(termText)

	return
}

func cofficientPart(termText string) (num float64, err error) {

	num, err = numberPart(termText)

	if err != nil {
		return
	}

	sign := signPart(termText)

	if sign == "-" {
		num = num * -1
	}

	return
}

func variablePart(termText string) (variable string) {

	variableRegExp := regexp.MustCompile("[A-z]")

	variable = variableRegExp.FindString(termText)

	return
}

func toTerm(termText string) (term Term, err error) {

	err = isValidTerm(termText)

	if err != nil {
		return
	}

	term.variable = variablePart(termText)

	term.coefficient, err = cofficientPart(termText)

	return
}


// consider variables of the form xar1