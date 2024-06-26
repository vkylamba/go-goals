package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"

	"go_goals/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Goal)
// DB Table: Plural (goals)
// Resource: Plural (Goals)
// Path: Plural (/goals)
// View Template Folder: Plural (/templates/goals/)

// GoalsResource is the resource for the Goal model
type GoalsResource struct {
	buffalo.Resource
}

// Selectable allows any struct to become an option in the select tag.
type PriorityOptions map[string]int

var PRIORITY_OPTIONS = PriorityOptions{
	"Low":    2,
	"Medium": 1,
	"High":   0,
}

var PRIORITY_IDS_TO_NAME = map[int]string{
	2: "Low",
	1: "Medium",
	0: "High",
}

func setPageContextForGoals(c buffalo.Context) {
	c.Set("priorityOptions", PRIORITY_OPTIONS)
	c.Set("priorityIdsToNameMapping", PRIORITY_IDS_TO_NAME)
	userOptions := make(map[string]string)
	userIdsToNameMap := make(map[string]string)
	users := []models.User{}
	err := models.DB.All(&users)
	if err == nil {
		for _, user := range users {
			userOptions[user.Name] = user.ID.String()
			userIdsToNameMap[user.ID.String()] = user.Name
		}
	} else {
		c.Logger().Errorf("Error getting users: %v", err)
	}
	c.Set("userOptions", userOptions)
	c.Set("userIdsToNameMap", userIdsToNameMap)
}

// List gets all Goals. This function is mapped to the path
// GET /goals
func (v GoalsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	goals := &models.Goals{}
	setPageContextForGoals(c)

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Goals from the DB
	if err := q.All(goals); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("goals", goals)
		return c.Render(http.StatusOK, r.HTML("goals/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		responseData := make(map[string]interface{})
		responseData["results"] = goals
		responseData["total"] = q.Paginator.TotalEntriesSize
		return c.Render(200, r.JSON(responseData))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(goals))
	}).Respond(c)
}

// Show gets the data for one Goal. This function is mapped to
// the path GET /goals/{goal_id}
func (v GoalsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Goal
	goal := &models.Goal{}
	setPageContextForGoals(c)

	// To find the Goal the parameter goal_id is used.
	if err := tx.Find(goal, c.Param("goal_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("goal", goal)

		return c.Render(http.StatusOK, r.HTML("goals/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(goal))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(goal))
	}).Respond(c)
}

// New renders the form for creating a new Goal.
// This function is mapped to the path GET /goals/new
func (v GoalsResource) New(c buffalo.Context) error {
	c.Set("goal", &models.Goal{})
	setPageContextForGoals(c)
	return c.Render(http.StatusOK, r.HTML("goals/new.plush.html"))
}

// Create adds a Goal to the DB. This function is mapped to the
// path POST /goals
func (v GoalsResource) Create(c buffalo.Context) error {
	// Allocate an empty Goal
	goal := &models.Goal{}
	setPageContextForGoals(c)

	// Bind goal to the html form elements
	if err := c.Bind(goal); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(goal)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("goal", goal)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("goals/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "goal.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/goals/%v", goal.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(goal))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(goal))
	}).Respond(c)
}

// Edit renders a edit form for a Goal. This function is
// mapped to the path GET /goals/{goal_id}/edit
func (v GoalsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Goal
	goal := &models.Goal{}
	setPageContextForGoals(c)

	if err := tx.Find(goal, c.Param("goal_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("goal", goal)
	return c.Render(http.StatusOK, r.HTML("goals/edit.plush.html"))
}

// Update changes a Goal in the DB. This function is mapped to
// the path PUT /goals/{goal_id}
func (v GoalsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Goal
	goal := &models.Goal{}
	setPageContextForGoals(c)

	if err := tx.Find(goal, c.Param("goal_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Goal to the html form elements
	if err := c.Bind(goal); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(goal)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("goal", goal)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("goals/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "goal.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/goals/%v", goal.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(goal))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(goal))
	}).Respond(c)
}

// Destroy deletes a Goal from the DB. This function is mapped
// to the path DELETE /goals/{goal_id}
func (v GoalsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Goal
	goal := &models.Goal{}
	setPageContextForGoals(c)

	// To find the Goal the parameter goal_id is used.
	if err := tx.Find(goal, c.Param("goal_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(goal); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "goal.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/goals")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(goal))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(goal))
	}).Respond(c)
}
