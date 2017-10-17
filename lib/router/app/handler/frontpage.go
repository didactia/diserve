package handler

import (
  "net/http"

  "diserve.didactia.org/lib/templater"
)

// Frontpage holds a pointer to the template, to be used for rendering the frontpage template.
type Frontpage struct {
  templater *templater.Templater
}

func (h *Frontpage) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  h.templater.Render(res, "frontpage", nil)
}

// NewFrontpage initializes, given a Templater, a new Frontpage handler.
func NewFrontpage(t *templater.Templater) *Frontpage {
  h := &Frontpage{
    templater: t,
  }
  return h
}
