package main

import (
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

func main() {
	e := new(Employee)
	http.HandleFunc("/", hello)
	http.HandleFunc("/employee", e.employeeInfo)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println("Webservice failed", err)
		return
	}
	fmt.Println("-----Webservice started-----")
}
