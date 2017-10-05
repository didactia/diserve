package handler

import (
  "net/http"
  "golang.org/x/text/language"
)

// Resource TODO
type Resource struct {
}

func (h *Resource) ServeHTTP(tag language.Tag) http.Handler {
  return http.HandlerFunc(func (res http.ResponseWriter, req *http.Request) {
    http.Error(res, "Not Implemented", http.StatusNotImplemented)
  })
}

// NewResource TODO
func NewResource() *Resource {
  h := &Resource{}
  return h
}
