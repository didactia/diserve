package env
import (
  "diserve.didactia.org/lib/util"
)

type variables struct {
  PORT string
  DBIP string
  DBPORT string
  STYLEPATH string
  TMPLPATH string
  PRELOADSTYLES bool
  LOADTMPLONREQUEST bool
}

var Vars variables

func Init(args []string) []string {
  head, tail := util.Shift(args)
  switch head {
  case "dev":
    development()
  default:
    development()
    return args
  }
  return tail
}
