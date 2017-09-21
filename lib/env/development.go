package env

func development() {
  Vars.PORT = "8000"
  Vars.DBIP = "localhost"
  Vars.DBPORT = "9080"
  Vars.STYLEDIR = "lib/frontend/style/"
}
