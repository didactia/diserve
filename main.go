package main

import (
  "os"
  "diserve.didactia.org/lib/router/args"
)

func main() {
  args.Args(os.Args[1:])
}
