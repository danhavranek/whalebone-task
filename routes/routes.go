package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/danhavranek/whalebone-task/models"
	"github.com/danhavranek/whalebone-task/repositories"
	"github.com/google/uuid"
)

func savePerson(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var newPerson models.PersonDTO
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newPerson)

	if err != nil {
		http.Error(w, "unable to decode json data", http.StatusBadRequest)
		return
	}
	// Values validation
	var newPersonUuid uuid.UUID
	newPersonUuid, err = uuid.Parse(newPerson.ExternalId)
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}
	_, err = mail.ParseAddress(newPerson.Email)
	if err != nil {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}
	var newPersonDateOfBirth time.Time
	newPersonDateOfBirth, err = time.Parse(time.RFC3339, newPerson.DateOfBirth)
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		return
	}

	// Create Person from PersonDTO
	personToBeStored := models.Person{ExternalId: newPersonUuid, Name: newPerson.Name, Email: newPerson.Email, DateOfBirth: models.CustomRFC3339Time{newPersonDateOfBirth}}

	// Check if person with given id already exists
	_, err = repositories.GetPersonById(personToBeStored.ExternalId)
	if err == nil {
		http.Error(w, "user with given id already exists", http.StatusConflict)
		return
	}

	err = repositories.CreatePerson(&personToBeStored)
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
	// Values validation
	parsedUuid, err := uuid.Parse(strings.Trim(req.URL.Path, "/"))
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	person, err := repositories.GetPersonById(parsedUuid)
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
