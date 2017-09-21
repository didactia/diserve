package server

import (
  "net/http"
  "fmt"
  "diserve.didactia.org/lib/env"
  "diserve.didactia.org/lib/router/app"
)

func Run(args []string) {
  app := app.NewApp()
  http.ListenAndServe(fmt.Sprintf(":%s", env.Vars.PORT), app)
}
