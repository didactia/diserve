package handler

import (
  "net/http"
  "diserve.didactia.org/lib/util"
  "diserve.didactia.org/lib/router/app/handler/api"
)

// API TODO
type API struct {
  User *api.User
  Concept *api.Concept
}

func (h *API) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  head, tail := util.ShiftPath(req.URL.Path)
  switch head {
  case "user":
    req.URL.Path = tail
    h.User.ServeHTTP(res, req)
  case "concept":
    req.URL.Path = tail
    h.Concept.ServeHTTP(res, req)
  default:
    http.Error(res, "Not Implemented", http.StatusNotImplemented)
  }
}

// NewAPI TODO
func NewAPI() *API {
  h := &API{
    User: api.NewUser(),
    Concept: api.NewConcept(),
  }
  return h
}
