package models

type Person struct {
	// TODO: add internal id?
	// Id          uint   `json:"-" gorm:"primaryKey"`
	ExternalId  string `json:"external_id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}
