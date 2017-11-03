package api

import (
  "net/http"
  "fmt"
  "diserve.didactia.org/lib/util"
  "diserve.didactia.org/lib/db"
  "diserve.didactia.org/lib/env"

  "github.com/oxtoacart/bpool"
)

// Concept TODO
type Concept struct {
  bufpool *bpool.BufferPool
  dbc *db.DatabaseClient
}

func (h *Concept) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  head, _ := util.ShiftPath(req.URL.Path)
  switch req.Method {
  case "POST":
    switch head {
    case "new":
      _, err := db.NewConcept("Testing", nil)
      if err != nil {
        http.Error(res, fmt.Sprint(err), http.StatusInternalServerError)
      }
    }
  default:
    http.Error(res, "Not Implemented", http.StatusNotImplemented)
  }
}

// NewConcept TODO
func NewConcept() *Concept {
  h := &Concept{
    bufpool: bpool.NewBufferPool(env.Vars.APICONCEPTBPOOLSIZE),
  }
  return h
}
