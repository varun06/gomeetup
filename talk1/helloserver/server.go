package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving", r.URL)
	fmt.Fprintln(w, "Hello chiGophers!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("serving on http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
