package env

func development() {
  Vars.PORT = "8000"
  Vars.DBIP = "localhost"
  Vars.DBPORT = "9080"
  Vars.STYLEPATH = "lib/frontend/style/*.css"
  Vars.TMPLPATH = "lib/frontend/tmpl/*.tmpl"
}
