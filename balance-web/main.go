package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/lancatlin/balance"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	equation := r.FormValue("a")
	if equation == "" {
		tpl.Execute(w, nil)
		return
	}
	type result struct {
		Q string
		A string
	}
	rs := result{equation, balance.Balance(equation)}
	err := tpl.Execute(w, rs)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Server start at http://localhost:8080")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
