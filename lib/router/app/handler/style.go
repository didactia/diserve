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
  styles map[string][]byte
}

func (h *Style) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  head, _ := util.ShiftPath(req.URL.Path)
  var data []byte
  var ok bool
  if env.Vars.PRELOADSTYLES {
    data, ok = h.styles[head]
  } else {
    data, ok = getStyle(head)
  }
  if ok {
    res.Header().Set("Content-Type", "text/css; charset=utf-8")
    res.Write(data)
  } else {
    http.Error(res, "Not Found", http.StatusNotFound)
  }
}

func NewStyle() *Style {
  h := &Style{
    styles: nil,
  }
  if env.Vars.PRELOADSTYLES {
    h.styles = getStyles()
  }
  return h
}

func getStyles() map[string][]byte {
  styles := make(map[string][]byte)
  paths, err := filepath.Glob(env.Vars.STYLEPATH + "*.css")
  if err != nil {
    log.Fatal(err)
  }
  for _, path := range paths {
    filename := filepath.Base(path)
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
      log.Fatal(err)
    }
    styles[filename] = bytes
  }
  return styles
}

func getStyle(filename string) ([]byte, bool) {
  bytes, err := ioutil.ReadFile(filepath.Join(env.Vars.STYLEPATH, filename))
  if err != nil {
    return nil, false
  }
  return bytes, true
}
