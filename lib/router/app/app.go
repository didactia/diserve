package app

import (
  "net/http"
  "diserve.didactia.org/lib/router/app/handler"
  "diserve.didactia.org/lib/util"
)

type App struct {
  Language *handler.Language
  API *handler.API
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  head, tail := util.ShiftPath(req.URL.Path)
  switch head {
  case "api":
    req.URL.Path = tail
    h.API.ServeHTTP(res, req)
  default:
    h.Language.ServeHTTP(res, req)
  }
}
