package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	Name     string
	LastName string
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello %v", r.URL.Path[1:])
	if err != nil {
		log.Println("Error occurred", r.URL)
		return
	}
}

var html = "<table>" +
	"<tr>" +
	"<th>Name</th>" +
	"<th>Last Name</th>" +
	"</tr>" +
	"<tr>" +
	"<td>%v</td>" +
	"<td>%v</td>" +
	"</tr>" +
	"<table>"

func (e *Employee) employeeInfo(w http.ResponseWriter, r *http.Request) {
	e.Name = "Harun"
	e.LastName = "Ugur"
	err := r.ParseForm()
	if err != nil {
		log.Println("Error occurred", err)
		return
	}
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	_, err = fmt.Fprintf(w, html, e.Name, e.LastName)
	if err != nil {
		log.Println("Error occurred", err)
		return
	}
}
func (e Employee) employeeJson(w http.ResponseWriter, _ *http.Request) {
	e.Name = "Cem"
	e.LastName = "Adrian"

	output, err := json.Marshal(e)
	if err != nil {
		log.Println("Parsing Error", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-TOKEN", "custom-token-value")
	_, err = fmt.Fprintf(w, string(output))
	if err != nil {
		log.Println("Serving Error", err)
		return
	}
}

func main() {
	e := new(Employee)
	http.HandleFunc("/", hello)
	http.HandleFunc("/employee", e.employeeInfo)
	http.HandleFunc("/employeeJson", e.employeeJson)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println("Webservice failed", err)
		return
	}
	fmt.Println("-----Webservice started-----")
}
