package handlers

import (
	"fmt"
	"html/template"
	"os"
)

var (
	tmpl      *template.Template
	errorTmpl *template.Template
)

/*
init parses the "templates/index.html" and "templates/error.html" templates, ensuring they are loaded before the application starts.
*/
func init() {
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Printf("error parsing the index template: %s\n", err)
		os.Exit(1)
	}
	errorTmpl, err = template.ParseFiles("templates/error.html")
	if err != nil {
		fmt.Printf("error parsing the error template: %s\n", err)
		os.Exit(1)
	}
}

/*
GetTemplate returns the globally stored index template, allowing consistent reuse throughout the application.
*/
func GetTemplate() *template.Template {
	return tmpl
}

/*
GetErrorTemplate returns the globally stored error template, allowing consistent reuse throughout the application.
*/
func GetErrorTemplate() *template.Template {
	return errorTmpl
}
