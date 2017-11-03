package handler

import (
  "net/http"
  "fmt"
  "diserve.didactia.org/lib/util"
  "diserve.didactia.org/lib/db"
  "diserve.didactia.org/lib/env"
  "diserve.didactia.org/lib/templater"
  "diserve.didactia.org/lib/frontend/local"
  "github.com/oxtoacart/bpool"
)

// Concept TODO
type Concept struct {
  bufpool *bpool.BufferPool
  templater *templater.Templater
}

func (h *Concept) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  head, _ := util.ShiftPath(req.URL.Path)
  switch head {
  case "":
    req.ParseForm()
    query := req.FormValue("q")
    h.showSearchResults(res, query)
  }
}

func (h *Concept) showSearchResults(res http.ResponseWriter, query string) {
  if query == "" {
    http.Error(res, "Bad Request", http.StatusBadRequest)
    return
  }
  response, err := db.GetConcepts(query)
  if err != nil {
    fmt.Println(err)
    http.Error(res, "Server Error", http.StatusInternalServerError)
    return
  }
  data := struct {
    Strings *local.Strings
    Concepts []db.Concept
  }{
    local.Map["en-US"],
    response,
  }
  fmt.Println(data.Strings.Prompts.NoResultsFound)
  buf := h.bufpool.Get()
  defer h.bufpool.Put(buf)
  h.templater.RenderBuffer(buf, "show-search-results", data)
  buf.WriteTo(res)
  return
}

// NewConcept TODO
func NewConcept(t *templater.Templater) *Concept {
  h := &Concept{
    templater: t,
    bufpool: bpool.NewBufferPool(env.Vars.BPOOLSIZE),
  }
  return h
}
