package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/pop/v6/slices"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Task is used by pop to map your tasks database table to your go code.
type Task struct {
	ID                 uuid.UUID     `json:"id" db:"id"`
	GoalID             uuid.UUID     `json:"goal_id" db:"goal_id"`
	MilestoneID        *uuid.UUID    `json:"milestone_id" db:"milestone_id"`
	Title              string        `json:"title" db:"title"`
	Description        nulls.String  `json:"description" db:"description"`
	ContributionFactor nulls.Float32 `json:"contribution_factor" db:"contribution_factor"`
	CompletionFactor   nulls.Float32 `json:"completion_factor" db:"completion_factor"`
	TargetDate         nulls.Time    `json:"target_date" db:"target_date"`
	CompletionDate     nulls.Time    `json:"completion_date" db:"completion_date"`
	Priority           nulls.Int     `json:"priority" db:"priority"`
	Tags               slices.String `json:"tags" db:"tags"`
	Active             nulls.Bool    `json:"active" db:"active"`
	Public             nulls.Bool    `json:"public" db:"public"`
	CreatedAt          time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Task) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Tasks is not required by pop and may be deleted
type Tasks []Task

// String is not required by pop and may be deleted
func (t Tasks) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Task) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Title, Name: "Title"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Task) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Task) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (t *Task) GetForUser(q *pop.Query, goalIds, milestoneIds []interface{}) ([]Task, error) {
	tasks := make([]Task, 0)
	goalIdsStr := ""
	milestoneIdsStr := ""
	for _, i := range goalIds {
		goalIdsStr += fmt.Sprintf("'%v',", i)
	}
	goalIdsStr = "(" + goalIdsStr[:len(goalIdsStr)-1] + ")"

	for _, i := range milestoneIds {
		milestoneIdsStr += fmt.Sprintf("'%v',", i)
	}
	milestoneIdsStr = "(" + milestoneIdsStr[:len(milestoneIdsStr)-1] + ")"

	query := `
		select *
		from tasks
		where
			(
				goal_id in ` + goalIdsStr + `
				or
				goal_id in ` + milestoneIdsStr + `
			)
			and active = 1
			and completion_date is null
	`

	q.RawQuery(query)
	err := q.All(&tasks)
	return tasks, err
}
