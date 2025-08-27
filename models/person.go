package models

type Person struct {
	ExternalId  string `json:"external_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}
