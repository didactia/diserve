package env
import (
  "diserve.didactia.org/lib/util"
  "strings"
  "log"
  "fmt"
)

type variables struct {
  PORT string
  DBIP string
  DBPORT string
  STYLEPATH string
  HTMLTMPLPATH string
  DBTMPLPATH string
  TMPLBPOOLSIZE int
  PRELOADSTYLES bool
  LOADTMPLONREQUEST bool
  USERTOKENDURATION int64
  HMACSECRET []byte
  APIUSERBPOOLSIZE int
}

// Vars is a struct that holds the current environment variables, needs to be initialised with Init(args []string)
var Vars variables

// Init initializes the Vars struct with the environment variables, the first string in args defines the environment,
// the following should be of the form "VAR:VALUE" witch overrides VAR with VALUE for the environment.
// With no matching environment, Vars will be initialized with development variables.
func Init(args []string) []string {
  head, tail := util.Shift(args)
  switch head {
  case "dev":
    development()
  default:
    development()
    return args
  }
  fmt.Println(Vars)
  tail, err := override(tail)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(Vars)
  return tail
}

func override(args []string) ([]string, error) {
  head, tail := util.Shift(args)
  fieldAndValue := strings.Split(head, ":")
  if len(fieldAndValue) != 2 {
    return args, nil
  }
  field := fieldAndValue[0]
  value := fieldAndValue[1]
  fieldType, err := util.FieldType(&Vars, field)
  if err != nil {
    return nil, err
  }
  typedValue, err := util.StringToType(value, fieldType)
  if err != nil {
    return nil, err
  }
  err = util.Override(&Vars, field, typedValue)
  if err != nil {
    return nil, err
  }
  return override(tail)
}
