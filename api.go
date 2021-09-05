package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	_, err := fmt.Fprintf(w, "Hello %v", r.URL.Path[1:])
	if err != nil {
		log.Println("Error occurred", r.URL)
		return
	}
}

func main()  {

	http.HandleFunc("/",hello)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println("Webservice failed",err)
		return
	}
	fmt.Println("-----Webservice started-----")
}
