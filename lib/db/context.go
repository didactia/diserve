package db

import (
  "errors"
  "log"
)

// ErrContextAlreadyExists self explanatory
var ErrContextAlreadyExists = errors.New("context already exists")

// ErrEmptyLabel is returned when trying to query for empty label
var ErrEmptyLabel = errors.New("no label given")

// Context TODO
type Context struct {
  Label string `json:"label"`
  UID uint64 `json:"_uid_"`
  UIDString string
}

type newContextQuery struct {
  Label string
}

// NewContext TODO
func NewContext(label string) (*Context, error) {
  context, err := FindContext(label)
  if err != nil {
    return context, ErrContextAlreadyExists
  }
  data := newContextQuery{
    Label: label,
  }
  res, err := Query("new-context", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  context = &Context{
    Label: label,
    UID: res.AssignedUids["context"],
    UIDString: uidString(res.AssignedUids["context"]),
  }
  return context, nil
}

type findContextResult struct {
  Context *Context `json:"context"`
}

type findContextQuery struct {
  Label string
}

// FindContext TODO
func FindContext(label string) (*Context, error) {
  if label == "" {
    return nil, ErrEmptyLabel
  }
  data := findContextQuery{
    Label: label,
  }
  var res findContextResult
  err := QueryAndUnmarshal("find-context", data, &res)
  if err != nil {
    return nil, err
  }
  return res.Context, nil
}
