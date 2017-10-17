package api

import (
  "time"
  "net/http"
  "diserve.didactia.org/lib/util"
  "diserve.didactia.org/lib/db"
  "diserve.didactia.org/lib/env"
  "diserve.didactia.org/lib/security"
  "encoding/json"
  "fmt"

  "github.com/oxtoacart/bpool"
)

// User TODO
type User struct {
  bufpool *bpool.BufferPool
  dbc *db.DatabaseClient
}

type loginResponse struct {
  Check bool `json:"c"`
  Err string `json:"e,omitempty"`
}

func (h *User) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  head, _ := util.ShiftPath(req.URL.Path)
  switch req.Method {
  case "POST":
    switch head {
    case "login":
      req.ParseForm()
      user, err := h.dbc.LoginUser(req.FormValue("u"), req.FormValue("p"))
      resp := loginResponse{}
      if err != nil {
        resp.Check = false
        resp.Err = fmt.Sprint(err)
      } else {
        token, exp, err := security.LoginToken(user)
        if err != nil {
          resp.Check = false
          resp.Err = fmt.Sprint(err)
        } else {
          resp.Check = true
          cookie := &http.Cookie{
            Name: "jwt",
            Value: token,
            Expires: time.Unix(exp, 0),
          }
          http.SetCookie(res, cookie)
        }
      }
      data, err := json.Marshal(resp)
      if err != nil {
        http.Error(res, fmt.Sprint(err), http.StatusInternalServerError)
      } else {
        res.Header().Set("Content-Type", "text/json; charset=utf-8")
        buf := h.bufpool.Get()
        defer h.bufpool.Put(buf)
        buf.Write(data)
        buf.WriteTo(res)
      }
    case "register":
      // TODO
      http.Error(res, "Not Implemented", http.StatusNotImplemented)
    default:
      http.Error(res, "Not Implemented", http.StatusNotImplemented)
    }
  default:
    http.Error(res, "Not Implemented", http.StatusNotImplemented)
  }
}

// NewUser TODO
func NewUser(dbc *db.DatabaseClient) *User {
  h := &User{
    bufpool: bpool.NewBufferPool(env.Vars.APIUSERBPOOLSIZE),
    dbc: dbc,
  }
  return h
}
