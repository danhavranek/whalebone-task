package models

// Since the service doesn't work with the data internally,
// all fields of the Person struct are stored as strings.
// Validation is handled in ../routes/routes.go.
// This decision allows us to preserve the original RFC3339
// time offset format without having to create methods
// to solve this problem.

type Person struct {
	ExternalId  string `json:"external_id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}
