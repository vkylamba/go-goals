package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/pop/v6/slices"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Goal is used by pop to map your goals database table to your go code.
type Goal struct {
	ID          uuid.UUID     `json:"id" db:"id"`
	User        User          `json:"user" db:"user"`
	Title       string        `json:"title" db:"title"`
	Description nulls.String  `json:"description" db:"description"`
	TargetDate  nulls.Time    `json:"target_date" db:"target_date"`
	Priority    nulls.Int     `json:"priority" db:"priority"`
	Tags        slices.String `json:"tags" db:"tags"`
	Active      nulls.Bool    `json:"active" db:"active"`
	Public      nulls.Bool    `json:"public" db:"public"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (g Goal) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Goals is not required by pop and may be deleted
type Goals []Goal

// String is not required by pop and may be deleted
func (g Goals) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (g *Goal) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: g.Title, Name: "Title"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (g *Goal) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (g *Goal) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
