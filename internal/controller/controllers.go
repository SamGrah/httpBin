package controllers

import "net/http"

type Controllers struct{}

type Deps struct {
	Index *http.HandlerFunc
}

func NewControllers(deps *Deps) *Controllers {
	return &Controllers{}
}

func (c *Controllers) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
