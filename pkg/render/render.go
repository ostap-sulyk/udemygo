package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const (
	templatesDir   = "./templates"
	layoutFileName = "base.layout.tmpl"
)

// RenderTemplate renders a template with the given name.
func RenderTemplaTE(w http.ResponseWriter, templ string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+templ, "./templates/base.layout.tmpl")

	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if we ahve template in our cache
	_, inMap := tc[t]
	if !inMap {
		// need to create a template
		log.Println("creating template and saving it to cache")
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println("Error creating template cache:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else {
		// we have template in our cache
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func createTemplateCache(tName string) error {
    tPath := filepath.Join(templatesDir, tName)
    layout := filepath.Join(templatesDir, layoutFileName)

    if tmpl, err := template.ParseFiles(tPath, layout); err != nil {
        return fmt.Errorf("template %q: %w", tName, err)
    } else {
        tc[tName] = tmpl
    }

    return nil
}
