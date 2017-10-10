package env

func development() {
  Vars.PORT = "8000"
  Vars.DBIP = "localhost"
  Vars.DBPORT = "9080"
  Vars.STYLEPATH = "lib/frontend/style/"
  Vars.HTMLTMPLPATH = "lib/frontend/tmpl/*.tmpl"
  Vars.PRELOADSTYLES = false
  Vars.LOADTMPLONREQUEST = true
}
