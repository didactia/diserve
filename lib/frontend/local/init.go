package local

import (
  "encoding/json"
  "diserve.didactia.org/lib/env"
  "io/ioutil"
  "log"
  "strings"
  "path/filepath"
)

// Map of local Strings structs, queriable by BCP47 tags.
var Map map[string]*Strings

// Init initializes the local Strings map.
func Init() {
  Map = make(map[string]*Strings)
  files, err := ioutil.ReadDir(env.Vars.LOCALSTRINGSPATH)
  if err != nil {
    log.Fatal(err)
  }
  for _, file := range files {
    basename := file.Name()
    name := strings.TrimSuffix(basename, filepath.Ext(basename))
    strs := NewStrings()
    data, err := ioutil.ReadFile(filepath.Join(env.Vars.LOCALSTRINGSPATH, basename))
    if err != nil {
      log.Fatal(err)
    }
    json.Unmarshal(data, strs)
    Map[name] = strs
  }
}
