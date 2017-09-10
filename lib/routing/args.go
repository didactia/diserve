package routing

import (
  "fmt"
  "log"
  "os"
  "diserve.didactia.org/lib/db"
  "diserve.didactia.org/lib/help"
  "diserve.didactia.org/lib/util"
)

// Args is a method for parsing and routing the arguments
func Args(args []string) {
  arg, args := util.Shift(os.Args[1:])
  switch arg {
  case "db":
    dbrouter(args)
  case "help":
    helprouter(args)
  case "run":
    runrouter(args)
  case "":
    fmt.Println("No argument given, use \"diserve help\" for a list of commands")
  default:
    fmt.Printf("No such command: %s, use \"diserve help\" for a list of commands\n", arg)
  }
}

// dbrouter will, given a slice of arguments, route to the appropriate db functionality
func dbrouter(args []string) {
  arg, args := util.Shift(args)
  switch arg {
  case "init":
    db.Init(args)
  default:
    fmt.Println("No argument given for command \"diserve db\", use \"diserve help db\" for a list of commands")
  }
}

// helprouter is a wrapper for the help parser
func helprouter(args []string) {
  help.Parse(args)
}

// runrouter will, given a slice of arguments, route to the appropriate run functionality
func runrouter(args []string) {
  log.Fatal("Not implemented")
}
