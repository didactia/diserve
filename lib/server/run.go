package server

import (
  "net/http"
  "fmt"
  "diserve.didactia.org/lib/env"
  "diserve.didactia.org/lib/db"
  "diserve.didactia.org/lib/router/app"
)

// Run runs the didactia server on the port PORT in the environment variables
func Run(args []string) {
  dbc := db.NewDatabaseClient(env.Vars.DBIP, env.Vars.DBPORT)
  defer dbc.Close()
  app := app.NewApp(dbc)
  http.ListenAndServe(fmt.Sprintf(":%s", env.Vars.PORT), app)
}
