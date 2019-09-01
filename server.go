package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"

	"solve"

	"github.com/rs/cors"
)

func solveHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("tmpl.html")

	res, err := solve.SolveStr(r.FormValue("eqns"), r.FormValue("tol"))
	if err != nil {
		return
	}

	t.ExecuteTemplate(w, "tbo", "")
	t.ExecuteTemplate(w, "th", res.Variables)
	t.ExecuteTemplate(w, "tr", res.Tbl)
	t.ExecuteTemplate(w, "tbc", "")
}

func validate(w http.ResponseWriter, r *http.Request) {

	acceptedCharactersExp := "(\\d|[A-z]|\\-|\\+|\\s|=)"

	acRegExp := regexp.MustCompile(acceptedCharactersExp)

	for _, ch := range r.FormValue("eqns") {

		if !acRegExp.MatchString(string(ch)) {

			errmg := "'" + string(ch) + "' is not accepted as a number, a symbol, variable or as part of one.\nPlease remove or change it to continue."

			fmt.Fprintf(w, errmg)

			return
		}
	}

	return
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}
	http.Handle("/solve", cors.AllowAll().Handler(http.HandlerFunc(solveHandler)))
	http.Handle("/validate", cors.AllowAll().Handler(http.HandlerFunc(validate)))
	http.Handle("/", cors.AllowAll().Handler(http.HandlerFunc(index)))

	server.ListenAndServe()
}
