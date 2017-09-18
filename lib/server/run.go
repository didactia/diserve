package server

import (
  "net/http"
  "fmt"
  "diserve.didactia.org/lib/env"
  "diserve.didactia.org/lib/router/app"
  "diserve.didactia.org/lib/router/app/handler"
)

func Run(args []string) {
  app := &app.App{
    Language: new(handler.Language),
    API: new(handler.API),
  }
  http.ListenAndServe(fmt.Sprintf(":%s", env.Vars.PORT), app)
}
