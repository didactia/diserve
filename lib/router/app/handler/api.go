package handler

import (
  "net/http"
)

type API struct {
}

func (h *API) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  http.Error(res, "Not Implemented", http.StatusNotImplemented)
  return
}
