package eqns

import (
	"regexp"
	"strings"
)

type Equation struct {
	variableTerms []Term
	constant      float64
}

func variableTermsPart(eqnText string) (terms []Term, err error) {

	termRegExp := regexp.MustCompile(varTermExp)

	termsStr := termRegExp.FindAllString(eqnText, -1)

	for _, termStr := range termsStr {

		term := Term{}

		term, err = toTerm(termStr)

		if err != nil {
			return
		}

		terms = append(terms, term)
	}

	return
}

func removeVariableTerms(eqnText string) (editted string) {

	editted = eqnText

	termRegExp := regexp.MustCompile(varTermExp)

	termsStr := termRegExp.FindAllString(eqnText, -1)

	for _, termStr := range termsStr {

		editted = strings.Replace(editted, termStr, "", 1)
	}

	return
}

func constantPart(eqnText string) (constant float64, err error) {

	constsStr := removeVariableTerms(eqnText)

	constRegExp := regexp.MustCompile(constExp)

	constsStrList := constRegExp.FindAllString(constsStr, -1)

	for _, constStr := range constsStrList {

		var newConstant float64

		newConstant, err = cofficientPart(constStr)

		if err != nil {
			return
		}

		constant = constant + newConstant
	}

	return
}

func ToEqn(eqnText string) (eqn Equation, err error) {

	eqn.variableTerms, err = variableTermsPart(eqnText)

	if err != nil {
		return
	}

	eqn.constant, err = constantPart(eqnText)

	if err != nil {
		return
	}

	return
}
