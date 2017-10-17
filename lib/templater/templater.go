package templater

import (
  "fmt"
  "log"
  "net/http"
  "text/template"

  "github.com/oxtoacart/bpool"

  "diserve.didactia.org/lib/env"
)

// Templater struct to hold the loaded templates.
type Templater struct {
  templates *template.Template
  bufpool *bpool.BufferPool
  path string
}

// Render will render the template of the given name, with the given data, to the given response writer.
// Will render an error message to the http.ResponseWriter if template is not found.
// Setting LOADTMPLONREQUEST to true in the environment variables will cause templates to be reloaded on all template renders.
func (t *Templater) Render(w http.ResponseWriter, name string, data interface{}) {
  var err error
  if env.Vars.LOADTMPLONREQUEST {
    t.templates, err = template.ParseGlob(t.path)
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

// RenderString renders template with name and data, and returns the rendered string.
func (t *Templater) RenderString(name string, data interface{}) (string, error) {
  var err error
  if env.Vars.LOADTMPLONREQUEST {
    t.templates, err = template.ParseGlob(t.path)
    if err != nil {
      return "", err
    }
  }
  tmpl := t.templates.Lookup(name)
  if tmpl == nil {
    return "", fmt.Errorf("the template %s does not exist", name)
  }
  buf := t.bufpool.Get()
  tmpl.Execute(buf, data)
  str := buf.String()
  t.bufpool.Put(buf)
  return str, nil
}

// NewTemplater returns a Templater with parsed templates in the directory TMPLPATH.
func NewTemplater(path string) *Templater {
  tmpls, err := template.ParseGlob(path)
  if err != nil {
    log.Fatal(err)
  }
  bufpool := bpool.NewBufferPool(env.Vars.TMPLBPOOLSIZE)
  templater := &Templater{
    templates: tmpls,
    bufpool: bufpool,
    path: path,
  }
  return templater
}
