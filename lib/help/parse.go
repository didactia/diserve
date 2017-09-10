package help

import (
  "fmt"
  "diserve.didactia.org/lib/util"
)

// Parse parses the arguments given and prints the appropriate help file
func Parse(args []string) {
  arg, args := util.Shift(args)
  commands, ok := Strings[arg]
  if ok {
    if args == nil {
      fmt.Printf("The following commands are available for \"diserve %s\":\n", arg)
      for command, helptext := range commands {
        fmt.Printf("  %s\n    %s", command, helptext)
      }
      print("\n")
    } else {
      helptext, ok := commands[args[0]]
      if ok {
        fmt.Printf("diserve %s %s\n  %s\n", arg, args[0], helptext)
      } else {
        fmt.Printf("No command: %s, for argument \"diserve %s\"", arg, args[0])
      }
    }
  } else {
    fmt.Printf("No command: diserve %s", arg)
  }
}
