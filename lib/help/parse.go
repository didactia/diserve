package help

import (
  "fmt"
  "diserve.didactia.org/lib/util"
)

// Parse parses the arguments given and prints the appropriate help strings
func Parse(args []string) {
  arg, args := util.Shift(args)
  commands, ok := Strings[arg]
  if ok {
    if args == nil {
      fmt.Printf("These are the available arguments for diserve %s:\n\n", arg)
      for command, helptext := range commands {
        fmt.Printf("  %s\n    %s", command, helptext)
      }
      print("\n")
    } else {
      helptext, ok := commands[args[0]]
      if ok {
        fmt.Printf("%s %s\n  %s\n", arg, args[0], helptext)
      } else {
        fmt.Printf("No argument: %s, for command \"diserve %s\"", arg, args[0])
      }
    }
  } else {
    if arg == "" {
      fmt.Println("These are the available commands for diserve:\n")
      for command := range Strings {
        fmt.Printf("  %s\n", command);
      }
      fmt.Println("\nuse \"diserve help <command>\" for a list of arguments for the command")
    } else {
      fmt.Printf("No command: diserve %s\n", arg)
    }
  }
}
