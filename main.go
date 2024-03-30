package main

import (
	"net/http"
	"text/template"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", choose)
	http.HandleFunc("/add", add)
	
	http.ListenAndServe(":8080", nil)
}

func choose(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	}
	tpl.Execute(w, nil)
}






func operation(w http.ResponseWriter, r *http.Request, operations string) {
	var result 
	number1 := r.FormValue("number1")
	number2 := r.FormValue("number2")

	num1, err := 

	if operations == "+" {
	result = number1 + number2
	}

}


