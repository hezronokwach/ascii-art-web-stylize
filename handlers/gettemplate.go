package handlers

import (
	"fmt"
	"html/template"
	"os"
)

var temp *template.Template

/* 
init parses the "templates/index.html" template and exits if there's an error, ensuring the application starts with necessary resources.
*/
func init() {
	temp1, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Printf("error parsing the template %s", err)
		os.Exit(1)
	}
	temp = temp1
}

/* 
GetTemplate returns the globally stored template, allowing consistent reuse throughout the application.
*/
func GetTemplate() *template.Template {
	return temp
}
