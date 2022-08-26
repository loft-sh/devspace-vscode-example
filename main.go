package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	fmt.Println("Started server on :9000")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
