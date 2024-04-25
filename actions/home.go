package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
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
