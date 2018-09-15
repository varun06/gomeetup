package main

import (
	"fmt"
	"log"
	"net/http"
)

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://localhost:8081"+r.RequestURI, http.StatusMovedPermanently)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving", r.URL)
	fmt.Fprintf(w, "Hello chiGophers!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	go http.ListenAndServeTLS("localhost:8081", "../testcerts/cert.pem", "../testcerts/key.pem", nil)
	http.ListenAndServe("localhost:8080", http.HandlerFunc(redirectToHTTPS))
}
