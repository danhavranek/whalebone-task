package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func initializeRoutes() {
	http.HandleFunc("/hello", hello)
}

func main() {
	initializeRoutes()

	http.ListenAndServe(":8090", nil)
}
