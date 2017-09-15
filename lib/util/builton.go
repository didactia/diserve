package util

import (
  "path"
  "strings"
)

// Shift returns the head and tail of a string slice, given an empty list it will return ("", nil).
func Shift(slice []string) (head string, tail []string) {
  switch len(slice) {
  case 0:
    return "", nil
  case 1:
    return slice[0], nil
  default:
    return slice[0], slice[1:]
  }
}

func ShiftPath(p string) (head, tail string) {
    p = path.Clean("/" + p)
    i := strings.Index(p[1:], "/") + 1
    if i <= 0 {
        return p[1:], "/"
    }
    return p[1:i], p[i:]
}
