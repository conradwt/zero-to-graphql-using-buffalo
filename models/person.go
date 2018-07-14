package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Person struct {
	ID          int64       `json:"id" db:"id"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	FirstName   string      `json:"first_name" db:"first_name"`
	LastName    string      `json:"last_name" db:"last_name"`
	Email       string      `json:"email" db:"email"`
	Username    string      `json:"username" db:"username"`
	Friendships Friendships `has_many:"friendships"`
}

// String is not required by pop and may be deleted
func (p Person) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// People is not required by pop and may be deleted
type People []Person

// String is not required by pop and may be deleted
func (p People) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Person) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.FirstName, Name: "FirstName"},
		&validators.StringIsPresent{Field: p.LastName, Name: "LastName"},
		&validators.StringIsPresent{Field: p.Email, Name: "Email"},
		&validators.StringIsPresent{Field: p.Username, Name: "Username"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Person) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Person) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
