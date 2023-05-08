package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	name := r.URL.Query().Get("name")
	if name != "" {
		msg = "Hello " + name
	}
	fmt.Fprintln(w, msg)
}

func main() {
	portNumber := ":8000"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		portNumber = ":" + val
	}
	http.HandleFunc("/api/HttpTrigger1", helloHandler)

	fmt.Printf("Server running on port %v", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		fmt.Printf("Error: !server start")
	}
}
