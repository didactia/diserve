package env

func development() {
  Vars.PORT = "8000"
  Vars.DBIP = "localhost"
  Vars.DBPORT = "9080"
  Vars.TMPLBPOOLSIZE = 48
  Vars.STYLEPATH = "lib/frontend/style/"
  Vars.HTMLTMPLPATH = "lib/frontend/tmpl/*.tmpl"
  Vars.DBTMPLPATH = "lib/db/tmpl/*.tmpl"
  Vars.PRELOADSTYLES = false
  Vars.LOADTMPLONREQUEST = true
  Vars.USERTOKENDURATION = 7 * 24 * 60 * 60
  Vars.HMACSECRET = []byte("not secret at all")
  Vars.APIUSERBPOOLSIZE = 48
}
