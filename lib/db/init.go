package db

import (
  "diserve.didactia.org/lib/env"
)

// Init initializes the database, with the schema defined at TODO
func Init(args []string) {
  InitializeDatabaseClient(env.Vars.DBIP, env.Vars.DBPORT)
  AddSchema(`name: string @index(exact) .
             password: password .
             title: string @index(exact, trigram) .
             label: string @index(exact, trigram) .
             prerequisite: uid .
             concept: uid .
             context: uid @reverse .
             perspective: uid .
             statement: uid .
             understander: uid @count .
             text: string .
             reasoning: uid .
             comment: uid .
             old: uid .
             next: uid .
             rating: uid @count .
             expression: uid .
             response: uid .`)
  dbc.Close()
}
