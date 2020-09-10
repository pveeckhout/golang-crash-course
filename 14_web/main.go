package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Server starting")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}