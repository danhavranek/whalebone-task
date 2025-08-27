package main

import (
	"fmt"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello\n")
}

func getPerson(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "%s\n", strings.Trim(req.URL.Path, "/"))
}

func initializeRoutes() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", getPerson)
}

func main() {
	initializeRoutes()

	http.ListenAndServe(":8090", nil)
}
