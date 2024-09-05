package actions

import (
	"fmt"
	"go_goals/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

// HomeHandler is a default handler to serve up
// a home page.
func RouteDetailsHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/routes.plush.html"))
}

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	time_window := "today"
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	goalIds, milestoneIds := SetPageContextForTasks(c)
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	var task models.Task
	tasks, err := task.GetForUser(q, goalIds, milestoneIds)
	if err != nil {
		return err
	}
	c.Set("time_window", time_window)
	c.Set("pagination", q.Paginator)
	c.Set("tasks", tasks)
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

func TokenAuthHandler(c buffalo.Context) error {
	token := "SAMPLE TOKEN"
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"token": token,
	}))
}

func ResourceHandler(c buffalo.Context) error {
	resourceData := map[string]interface{}{
		"baseUrl": "http://localhost:3000/api",
		"resources": []map[string]interface{}{
			{
				"name":            "goals",
				"label":           "Goals",
				"title":           "Available goals",
				"icon":            "AddHomeWork",
				"detailsFormType": "simple",
				"allowedActions": []string{
					"filterItems",
					"selectColumns",
					"showItem",
					"editItem",
					"createItem",
					"exportItems",
				},
				"fields": []map[string]interface{}{
					{
						"source":       "id",
						"label":        "ID",
						"type":         "text",
						"showInList":   true,
						"editable":     false,
						"size":         2,
						"formPosition": []int{1, 1},
					},
					{
						"source":       "name",
						"label":        "Name",
						"type":         "text",
						"showInList":   true,
						"editable":     true,
						"size":         10,
						"formPosition": []int{1, 2},
					},
					{
						"source":       "avatar",
						"label":        "Avatar",
						"type":         "avatar",
						"showInList":   true,
						"editable":     true,
						"size":         12,
						"formPosition": []int{2, 1},
					},
				},
				"filters": []map[string]interface{}{
					{
						"label": "Project type",
						"field": "projectType",
						"values": []map[string]interface{}{
							{"label": "Test", "value": "test"},
							{"label": "Test - 1", "value": "test1"},
						},
					},
				},
			},
		},
	}
	return c.Render(http.StatusOK, r.JSON(resourceData))
}
