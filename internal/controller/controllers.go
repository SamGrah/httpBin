package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type Controllers struct{}

type Deps struct{}

func NewControllers(deps *Deps) *Controllers {
	return &Controllers{}
}

func (c *Controllers) Index(w http.ResponseWriter, r *http.Request) {
	tmplFile := "index.html"

	t := template.Must(template.ParseFiles(tmplFile))

	w.Header().Set("Content-Type", "text/html")
	err := t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
