package templater

import (
  "fmt"
  "log"
  "net/http"
  "text/template"

  "diserve.didactia.org/lib/env"
)

// Templater struct to hold the loaded templates.
type Templater struct {
  templates *template.Template
}

// Render will render the template of the given name, with the given data, to the given response writer.
// Will render an error message to the http.ResponseWriter if template is not found.
// Setting LOADTMPLONREQUEST to true in the environment variables will cause templates to be reloaded on all template renders.
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

// NewTemplater returns a Templater with parsed templates in the directory TMPLPATH.
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
