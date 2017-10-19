package db

import (
  "errors"
  "log"
)

// ErrNoConceptFound is given when no concepts are found
var ErrNoConceptFound = errors.New("no concept found")

// ErrTooManyConcepts is given when too many concepts are found in a query
// expecting only one.
var ErrTooManyConcepts = errors.New("too many concepts")

// ErrConceptAlreadyExists is given on new concept creation, when a concept of
// that name and context already exists.
var ErrConceptAlreadyExists = errors.New("concept already exists")

// Concept TODO
type Concept struct {
  Title string `json:"title"`
  UID uint64 `json:"_uid_"`
  Context *Context `json:"context"`
  Perspectives []Perspective `json:"perspectives"`
}

type newConceptQuery struct {
  Title string
}

// NewConcept inserts a new concept into the database and returns the concept
// struct.
func NewConcept(title string, context *Context) (*Concept, error) {
  uid, err := FindConceptUID(title, context)
  if uid != 0 {
    return nil, ErrConceptAlreadyExists
  }
  if err != ErrNoConceptFound {
    return nil, err
  }
  data := newConceptQuery{
    Title: title,
  }
  result, err := Query("new-concept", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  concept := &Concept{
    Title: title,
    UID: result.AssignedUids["concept"],
  }
  return concept, nil
}

type findConceptUIDResponse struct {
  Concepts []Concept `json:"concept"`
}

type findConceptUIDQuery struct {
  Title string
  Context *Context
}

// FindConceptUID finds the UID of the concept with the given title in the
// given context, if context is nil, and there exists only one concept of the
// title, that concept UID is returned, if there exists more concepts an error
// is returned.
func FindConceptUID(title string, context *Context) (uint64, error) {
  data := findConceptUIDQuery{
    Title: title,
    Context: context,
  }
  var res findConceptUIDResponse
  err := QueryAndUnmarshal("find-concept-uid", data, &res)
  if err != nil {
    return 0, err
  }
  switch len(res.Concepts) {
  case 0:
    return 0, ErrNoConceptFound
  case 1:
    return res.Concepts[0].UID, nil
  default:
    return 0, ErrTooManyConcepts
  }
}

type getConceptsResponse struct {
  Concepts []Concept `json:"concepts"`
}

type getConceptsQuery struct {
  Title string
}

// GetConcepts will given a concept title return a slice of concepts with
// the given title, differing in context
func GetConcepts(title string) ([]Concept, error) {
  data := getConceptsQuery{
    Title: title,
  }
  var result getConceptsResponse
  err := QueryAndUnmarshal("get-concepts-by-title", data, &result)
  if err != nil {
    return nil, err
  }
  if len(result.Concepts) == 0 {
    return nil, ErrNoConceptFound
  }
  return result.Concepts, nil
}
