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

// Milestone is used by pop to map your milestones database table to your go code.
type Milestone struct {
	ID                 uuid.UUID     `json:"id" db:"id"`
	GoalID             uuid.UUID     `json:"goal_id" db:"goal_id"`
	Title              string        `json:"title" db:"title"`
	Description        nulls.String  `json:"description" db:"description"`
	ContributionFactor nulls.Float32 `json:"contributionFactor" db:"contribution_factor"`
	CompletionFactor   nulls.Float32 `json:"completionFactor" db:"completion_factor"`
	TargetDate         nulls.Time    `json:"targetDate" db:"target_date"`
	Priority           nulls.Int     `json:"priority" db:"priority"`
	Tags               slices.String `json:"tags" db:"tags"`
	Active             nulls.Bool    `json:"active" db:"active"`
	Public             nulls.Bool    `json:"public" db:"public"`
	CreatedAt          time.Time     `json:"createdAt" db:"created_at"`
	UpdatedAt          time.Time     `json:"updatedAt" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (m Milestone) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Milestones is not required by pop and may be deleted
type Milestones []Milestone

// String is not required by pop and may be deleted
func (m Milestones) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *Milestone) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: m.Title, Name: "Title"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *Milestone) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *Milestone) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
