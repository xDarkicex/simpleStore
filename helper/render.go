package helper

import (
	"fmt"
	"html/template"
)

var templates *template.Template

//Render fucntion renders page
func Render(a RouterArgs, object map[string]interface{}) {
	templates = template.Must(template.ParseGlob("./app/views/layout/*"))
	err := templates.ExecuteTemplate(a.Response, "index.gohtml", object)
	if err != nil {
		fmt.Print(err)
	}
}
