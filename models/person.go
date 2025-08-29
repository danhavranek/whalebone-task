package models

import (
	"github.com/google/uuid"
)

// Validation during conversion from from PersonDTO into Person
// is handled in ../routes/routes.go.

type Person struct {
	ExternalId  uuid.UUID         `json:"external_id" gorm:"primaryKey"`
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	DateOfBirth CustomRFC3339Time `json:"date_of_birth"`
}

type PersonDTO struct {
	ExternalId  string `json:"external_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}
