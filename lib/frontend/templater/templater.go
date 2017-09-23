package templater

import (
  "fmt"
  "log"
  "net/http"
  "text/template"

  "diserve.didactia.org/lib/env"
)

type Templater struct {
  templates *template.Template
}

func (t *Templater) Render(w http.ResponseWriter, name string, data interface{}) {
  var err error
  if env.Vars.LOADTMPLONREQUEST {
    t.templates, err = template.ParseGlob(env.Vars.TMPLPATH)
    if err != nil {
      http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
      return
    }
  }
  tmpl := t.templates.Lookup(name)
  if tmpl == nil {
    http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
      http.StatusInternalServerError)
    return
  }
  tmpl.Execute(w, data)
}

func NewTemplater() *Templater {
  tmpls, err := template.ParseGlob(env.Vars.TMPLPATH)
  if err != nil {
    log.Fatal(err)
  }
  templater := &Templater{
    templates: tmpls,
  }
  return templater
}
