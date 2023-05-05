package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "Error with parsing the form %v", err)
		return
	}

	fmt.Fprintln(res, "POST Request")
	name := req.FormValue("username")
	password := req.FormValue("password")

	fmt.Fprintf(res, "Name: %s \n", name)
	fmt.Fprintf(res, "Password: %s \n", password)
}
func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello GET")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting go on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
