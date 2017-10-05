package help

var db = map[string]string{
  "init":"initializes the database with schema definitions of TODO",
}

// Strings is a map of map of strings, containing the two first levels of commands and their help messages.
var Strings = map[string]map[string]string{
  "db":db,
}
