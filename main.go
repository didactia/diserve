package main

import (
  "os"
  "diserve.didactia.org/lib/routing"
)

func main() {
  routing.Args(os.Args[1:])
}
