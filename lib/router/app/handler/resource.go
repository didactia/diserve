package handler

import (
  "net/http"
  "golang.org/x/text/language"
)

type Resource struct {
}

func (h *Resource) Handler(tag language.Tag) http.Handler {
  return http.HandlerFunc(func (res http.ResponseWriter, req *http.Request) {
    http.Error(res, "Not Implemented", http.StatusNotImplemented)
  })
}

func NewResource() *Resource {
  h := &Resource{}
  return h
}
