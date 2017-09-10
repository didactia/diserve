package main

import (
  "os"
  routing "diserve.didactia.org/lib/routing"
)

func main() {
  routing.Args(os.Args[1:])
}
