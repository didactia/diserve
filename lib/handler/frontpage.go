package handler

import (
  "net/http"

  "diserve.didactia.org/lib/templater"
  "diserve.didactia.org/lib/env"

  "github.com/oxtoacart/bpool"
)

// Frontpage holds a pointer to the template, to be used for rendering the frontpage template.
type Frontpage struct {
  templater *templater.Templater
  bufpool *bpool.BufferPool
}

func (h *Frontpage) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  buf := h.bufpool.Get()
  defer h.bufpool.Put(buf)
  h.templater.RenderBuffer(buf, "frontpage", nil)
  buf.WriteTo(res)
}

// NewFrontpage initializes, given a Templater, a new Frontpage handler.
func NewFrontpage(t *templater.Templater) *Frontpage {
  h := &Frontpage{
    templater: t,
    bufpool: bpool.NewBufferPool(env.Vars.BPOOLSIZE),
  }
  return h
}
