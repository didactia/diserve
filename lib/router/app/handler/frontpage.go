package handler

import (
  "net/http"

  "diserve.didactia.org/lib/frontend/templater"
)

type Frontpage struct {
  templater *templater.Templater
}

func (h *Frontpage) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  h.templater.Render(res, "frontpage", nil)
}

func NewFrontpage(t *templater.Templater) *Frontpage {
  h := &Frontpage{
    templater: t,
  }
  return h
}
