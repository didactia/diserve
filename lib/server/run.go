package server

import (
  "net/http"
  "diserve.didactia.org/lib/router/app"
  "diserve.didactia.org/lib/router/app/handler"
)

func Run(args []string) {
  app := &app.App{
    Language: new(handler.Language),
    API: new(handler.API),
  }
  http.ListenAndServe(":8000", app)
}
