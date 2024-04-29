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

// Note is used by pop to map your notes database table to your go code.
type Note struct {
	ID          uuid.UUID     `json:"id" db:"id"`
	GoalID      nulls.UUID    `json:"goal_id" db:"goal_id"`
	MilestoneID nulls.UUID    `json:"milestone_id" db:"milestone_id"`
	TaskID      nulls.UUID    `json:"task_id" db:"task_id"`
	NoteID      nulls.UUID    `json:"note_id" db:"note_id"`
	CreatedBy   nulls.UUID    `json:"created_by" db:"created_by"`
	Title       string        `json:"title" db:"title"`
	Type        string        `json:"type" db:"type"`
	Description nulls.String  `json:"description" db:"description"`
	Tags        slices.String `json:"tags" db:"tags"`
	Active      nulls.Bool    `json:"active" db:"active"`
	Public      nulls.Bool    `json:"public" db:"public"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (n Note) String() string {
	jn, _ := json.Marshal(n)
	return string(jn)
}

// Notes is not required by pop and may be deleted
type Notes []Note

// String is not required by pop and may be deleted
func (n Notes) String() string {
	jn, _ := json.Marshal(n)
	return string(jn)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (n *Note) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: n.Title, Name: "Title"},
		&validators.StringIsPresent{Field: n.Type, Name: "Type"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (n *Note) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (n *Note) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
