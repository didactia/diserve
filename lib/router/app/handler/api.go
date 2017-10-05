package handler

import (
  "net/http"
)

// API TODO
type API struct {
}

func (h *API) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  http.Error(res, "Not Implemented", http.StatusNotImplemented)
  return
}

// NewAPI TODO
func NewAPI() *API {
  h := &API{}
  return h
}
