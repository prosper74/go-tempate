package handlers

import (
	"net/http"

	"github.com/atuprosper/go-project/pkg/config"
	"github.com/atuprosper/go-project/pkg/models"
	"github.com/atuprosper/go-project/pkg/render"
)

// Creating a Repository pattern
// This variable is the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// This function creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// This function NewHandlers, sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// A function with a 'reciever' m, of type 'Repository'. This will give our handler function access to everything in the config file
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	// Send the data to the template
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}
