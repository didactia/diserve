package handler

import (
  "net/http"
  "diserve.didactia.org/lib/util"
  "golang.org/x/text/language"
)

type Language struct {
  Resource *Resource
}

func (h *Language) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  var head string
  head, req.URL.Path = util.ShiftPath(req.URL.Path)
  tag, err := language.Parse(head)
  if err != nil {
    http.Error(res, "Not Found", http.StatusNotFound)
  } else {
    h.Resource.Handler(tag).ServeHTTP(res, req)
  }
}
