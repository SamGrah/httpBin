package controllers

import (
	"context"
	"log"
	"net/http"

	"app/internal/templates"
)

type Controllers struct{}

type Deps struct{}

func NewControllers(deps *Deps) *Controllers {
	return &Controllers{}
}

func (c *Controllers) Index(w http.ResponseWriter, r *http.Request) {
	component := templates.Hello("Fred")

	err := component.Render(context.Background(), w)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "text/html")
}
