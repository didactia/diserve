package db

import (
  "diserve.didactia.org/lib/env"
)

// Init initializes the database, with the schema defined at TODO
func Init(args []string) {
  dbc := NewDatabaseClient(env.Vars.DBIP, env.Vars.DBPORT)
  defer dbc.Close()
  dbc.AddSchema(`name: string @index(exact) .
                 password: password .
                 title: string @index(exact) .
                 prerequisite: uid .
                 concept: uid .
                 understander: uid @count .
                 text: string .
                 reasoning: uid .
                 comment: uid .
                 old: uid .
                 next: uid .
                 rating: uid @count .
                 expression: uid .
                 response: uid .`)
}
