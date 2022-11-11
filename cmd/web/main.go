package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atuprosper/go-project/pkg/config"
	"github.com/atuprosper/go-project/pkg/handlers"
	"github.com/atuprosper/go-project/pkg/render"
)

const port = ":8080"

// Building a web app
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	// Render the NewTemplates and add a reference to the AppConfig
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Server started at port %s", port))
	http.ListenAndServe(port, nil)
}
