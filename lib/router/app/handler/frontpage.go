package handler

import (
  "net/http"
)

type Frontpage struct {
}

func (h *Frontpage) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  http.Error(res, "Not Found", http.StatusNotFound)
}

func NewFrontpage() *Frontpage {
  h := &Frontpage{}
  return h
}
