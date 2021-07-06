package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello")
}
