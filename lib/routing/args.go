package routing

import (
  "fmt"
  "log"
  "os"
  "diserve.didactia.org/lib/db"
  "diserve.didactia.org/lib/help"
  "diserve.didactia.org/lib/util"
)

func Args(args []string) {
  arg, args := util.Shift(os.Args[1:])
  switch arg {
  case "db":
    _db(args)
  case "help":
    _help(args)
  case "run":
    _run(args)
  default:
    fmt.Println("No argument given, run \"diserve help\" for a list of commands")
  }
}

func _db(args []string) {
  arg, args := util.Shift(args)
  switch arg {
  case "init":
    db.Init(args)
  default:
    fmt.Println("No argument given for command \"diserve db\", run \"diserve help db\" for a list of commands")
  }
}

func _help(args []string) {
  help.Parse(args)
}

func _run(args []string) {
  log.Fatal("Not implemented")
}
