package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/atuprosper/go-project/pkg/config"
)

// A FuncMap is a map of functions that we can use in our template
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// This will rebuild the template cache on every page render. It is really not efficient. A better way is to set application wide configuration, that will build the template cache only ones and then use it for every render. To do this we create a new config file in pkg folder
	// tc, err := CreateTemplateCache()

	// Get the template cache from the app config
	tc := app.TemplateCache

	// If there are no template cache, die the server. If we get passed this, then we have our template cache
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// t holds the actual template found, while "ok" will return true if the template exists in our directory. If we get passed this, then we have the actual template that we want to use
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from the template cache")
	}

	// Create a bytesBuffer that will hold the information of the parsed template in memory, and put them in a byte
	buf := new(bytes.Buffer)

	//Execeute the tamplate file and put it in the buffer
	_ = t.Execute(buf, nil)

	// Write the buffer to the resposeWriter(browser)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// A [map] data structure to look up things very quickly. myCache is a cache that will hold all the templates
	myCache := map[string]*template.Template{}

	// get all the pages in the "templates directory" that ends with page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// Loop through all the pages to get individual page and the filepath of the page\
	// The underscore _ means we are ignoring the index of the list
	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)

		// Create a template set (ts), that will have functions "Funcs(functions), which are external functions not build into Go language"
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Look for any files that ends with layout.html in the templates directory
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// If we find any layout.html file, we want to pass them to our template set (ts)
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		// Add the template set(ts) we just created to our cache
		myCache[name] = ts
	}

	// Return myCache and ignore the value for error using nil. We have already dealt with all the posible errors
	return myCache, nil
}
