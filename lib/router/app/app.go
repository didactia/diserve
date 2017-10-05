package app

import (
  "net/http"
  "diserve.didactia.org/lib/router/app/handler"
  "diserve.didactia.org/lib/util"
  "diserve.didactia.org/lib/frontend/templater"
)

// App is the first level struct of the router, this contains the handlers for all root http requests.
type App struct {
  Language *handler.Language
  API *handler.API
  Style *handler.Style
  Frontpage *handler.Frontpage
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  head, tail := util.ShiftPath(req.URL.Path)
  switch head {
  case "api":
    req.URL.Path = tail
    h.API.ServeHTTP(res, req)
  case "style":
    req.URL.Path = tail
    h.Style.ServeHTTP(res, req)
  case "":
    req.URL.Path = tail //nil
    h.Frontpage.ServeHTTP(res, req)
  default:
    h.Language.ServeHTTP(res, req)
  }
}

// NewApp returns a new App with its handlers initialized.
func NewApp() *App {
  t := templater.NewTemplater()
  h := &App {
    Language: handler.NewLanguage(),
    API: handler.NewAPI(),
    Style: handler.NewStyle(),
    Frontpage: handler.NewFrontpage(t),
  }
  return h
}
