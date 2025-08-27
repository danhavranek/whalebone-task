package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func savePerson(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	data, _ := io.ReadAll(req.Body)
	fmt.Fprintf(w, "%s\n", data)
	req.Body.Close()
}

func getPerson(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "%s\n", strings.Trim(req.URL.Path, "/"))
}

func initializeRoutes() {
	http.HandleFunc("/save", savePerson)
	http.HandleFunc("/", getPerson)
}

func main() {
	initializeRoutes()

	http.ListenAndServe(":8090", nil)
}
