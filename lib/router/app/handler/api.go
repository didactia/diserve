package handler

import (
  "net/http"
  "diserve.didactia.org/lib/db"
  "diserve.didactia.org/lib/util"
  "diserve.didactia.org/lib/router/app/handler/api"
)

// API TODO
type API struct {
  User *api.User
}

func (h *API) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  head, tail := util.ShiftPath(req.URL.Path)
  switch head {
  case "user":
    req.URL.Path = tail
    h.User.ServeHTTP(res, req)
  default:
    http.Error(res, "Not Implemented", http.StatusNotImplemented)
  }
}

// NewAPI TODO
func NewAPI(dbc *db.DatabaseClient) *API {
  h := &API{
    User: api.NewUser(dbc),
  }
  return h
}
