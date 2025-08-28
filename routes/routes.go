package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/danhavranek/whalebone-task/models"
	"github.com/danhavranek/whalebone-task/repositories"
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
		http.Error(w, "unable to decode json data", http.StatusInternalServerError)
		return
	}
	err = repositories.CreatePerson(&newPerson)
	if err != nil {
		http.Error(w, "unable to store data into db", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getPerson(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	person, err := repositories.GetPersonById(strings.Trim(req.URL.Path, "/"))
	if err != nil {
		http.Error(w, "no user with given id found", http.StatusNotFound)
		return
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		http.Error(w, "error searializing data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", string(jsonData))
}

func Init() {
	http.HandleFunc("/save", savePerson)
	http.HandleFunc("/", getPerson)
}
