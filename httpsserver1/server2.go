package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving", r.URL)
	fmt.Fprintf(w, "Hello chiGophers!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	go http.ListenAndServeTLS("localhost:8081", "../testcerts/cert.pem", "../testcerts/key.pem", nil)

	http.ListenAndServe("localhost:8080", nil)
}
