package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/danhavranek/whalebone-task/models"
	"github.com/google/uuid"
)

func savePerson(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var newPerson models.Person
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newPerson)
	if err != nil {
		panic(err)
	}
	// TODO: save data into DB
	w.WriteHeader(http.StatusOK)
}

func getPerson(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// fmt.Fprintf(w, "%s\n", strings.Trim(req.URL.Path, "/"))
	// TODO: keep original offset in timestamp
	personMock := models.Person{ExternalId: uuid.NewString(), Name: "Test Person", Email: "test@example.com", DateOfBirth: time.Now().UTC().Format(time.RFC3339)}
	data, err := json.Marshal(personMock)
	if err != nil {
		// TODO: handle panic branches as server errors
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", string(data))
}

func initializeRoutes() {
	http.HandleFunc("/save", savePerson)
	http.HandleFunc("/", getPerson)
}

func main() {
	initializeRoutes()

	http.ListenAndServe(":8090", nil)
}
