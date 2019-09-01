package eqns

var signExp, spaceExp, numberExp, constExp, varTermExp, termExp, eqnExp string

func init() {

	signExp = "(\\-|\\+)?"

	spaceExp = "(\\s)*"

	numberExp = "((\\d)+(\\.)?(\\d)*)"

	constExp = signExp + spaceExp + numberExp

	varTermExp = constExp + "?[A-Za-z]{1}"

	termExp = spaceExp + "(" + constExp + "|" + varTermExp + ")" + spaceExp

	eqnExp = "(" + termExp + ")+=(" + termExp + ")+"
}
