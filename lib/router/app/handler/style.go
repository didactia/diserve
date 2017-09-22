package handler

import (
  "net/http"
  "io/ioutil"
  "path/filepath"
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
    res.Header().Set("Content-Type", "text/css; charset=utf-8")
    res.Write(data)
  }
}

func NewStyle() *Style {
  h := &Style{
    files: make(map[string][]byte),
  }
  paths, err := filepath.Glob(env.Vars.STYLEPATH)
  if err != nil {
    log.Fatal(err)
  }
  for _, path := range paths {
    filename := filepath.Base(path)
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
      log.Fatal(err)
    }
    h.files[filename] = bytes
  }
  return h
}
