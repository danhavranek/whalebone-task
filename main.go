package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello\n")
}

func initializeRoutes() {
	http.HandleFunc("/hello", hello)
}

func main() {
	initializeRoutes()

	http.ListenAndServe(":8090", nil)
}
