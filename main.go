package main

import (
  "os"
  "diserve.didactia.org/lib/routers"
)

func main() {
  routers.Args(os.Args[1:])
}
