package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"
	"github.com/pkg/errors"

	"go_goals/models"
)

// UsersResource is the resource for the User model
type UsersResource struct {
	buffalo.Resource
}

func (v UsersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	users := &models.Users{}
	setPageContextForGoals(c)

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Goals from the DB
	if err := q.All(users); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("users", users)
		return c.Render(http.StatusOK, r.HTML("users/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		responseData := make(map[string]interface{})
		responseData["results"] = users
		responseData["total"] = q.Paginator.TotalEntriesSize
		return c.Render(200, r.JSON(responseData))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(users))
	}).Respond(c)
}

func (v UsersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}
	setPageContextForGoals(c)

	// To find the user the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("user", user)

		return c.Render(http.StatusOK, r.HTML("users/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(user))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(user))
	}).Respond(c)
}

// Edit renders a edit form for a User. This function is
// mapped to the path GET /users/{user_id}/edit
func (v UsersResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}
	setPageContextForGoals(c)

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("users/edit.plush.html"))
}

func (v UsersResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Goal
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind User to the html form elements
	if err := c.Bind(user); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(user)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("user", user)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("users/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "user.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/users/%v", user.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(user))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(user))
	}).Respond(c)
}

// New renders the form for creating a new User.
// This function is mapped to the path GET /user/new
func (v UsersResource) New(c buffalo.Context) error {
	c.Set("user", &models.Goal{})
	return c.Render(http.StatusOK, r.HTML("users/new.plush.html"))
}

// UsersCreate registers a new user with the application.
func (v UsersResource) Create(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)
		return c.Render(http.StatusOK, r.HTML("users/new.plush.html"))
	}

	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome to go-goals!")

	return c.Redirect(http.StatusFound, "/")
}

// Destroy deletes a User from the DB. This function is mapped
// to the path DELETE /user/{user_id}
func (v UsersResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the user the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(user); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "user.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/users")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(user))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(user))
	}).Respond(c)
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				c.Logger().Warnf("user attempted to access with current_user_id '%v' that is not found: %v", uid, err)

				c.Session().Delete("current_user_id")
				c.Session().Set("redirectURL", c.Request().URL.String())
				c.Flash().Add("danger", "You must be authorized with a correct user to see that page")
				return c.Redirect(http.StatusFound, "/auth/new")
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Session().Set("redirectURL", c.Request().URL.String())

			err := c.Session().Save()
			if err != nil {
				return errors.WithStack(err)
			}

			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(http.StatusFound, "/auth/new")
		}
		return next(c)
	}
}
