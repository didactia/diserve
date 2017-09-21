package handler

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "log"

  "diserve.didactia.org/lib/util"
  "diserve.didactia.org/lib/env"
)

type Style struct {
  files map[string][]byte
}

func (h *Style) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  head, _ := util.ShiftPath(req.URL.Path)
  if data := h.files[head]; data == nil {
    http.Error(res, "Not Found", http.StatusNotFound)
  } else {
    res.Write(data)
  }
}

func NewStyle() *Style {
  h := &Style{
    files: make(map[string][]byte),
  }
  files, err := ioutil.ReadDir(env.Vars.STYLEDIR)
  if err != nil {
    log.Fatal(err)
  }
  for _, file := range files {
    filename := file.Name()
    bytes, err := ioutil.ReadFile(fmt.Sprintf("%s%s", env.Vars.STYLEDIR, filename))
    if err != nil {
      log.Fatal(err)
    }
    h.files[filename] = bytes
  }
  return h
}
