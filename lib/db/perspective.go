package db

import (
  "errors"
  "log"
)

// ErrNoPerspectivesFound is given when no perspectives are found
var ErrNoPerspectivesFound = errors.New("no perspectives found")

// ErrNoStatements is given on creating a new perspective, with a zero length
// statement slice.
var ErrNoStatements = errors.New("cannot create new perspective with no statements")

// Perspective TODO
type Perspective struct {
  UID uint64 `json:"_uid_"`
  UIDString string
  Statement *Statement `json:"statement"`
}

type newPerspectiveQuery struct {
  ConceptUID string
  IDAndTexts []idAndText
}

// NewPerspective inserts a new perspective into the database and returns the concept
// struct.
func (c *Concept) NewPerspective(texts []string) (*Perspective, error) {
  if len(texts) == 0 {
    return nil, ErrNoStatements
  }
  iAndT := identifyTexts(texts)
  data := newPerspectiveQuery{
    ConceptUID: uidString(c.UID),
    IDAndTexts: iAndT,
  }
  res, err := Query("new-perspective", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  statement, err := statementSerializer(iAndT, res.AssignedUids)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseUnmarshalling
  }
  perspective := &Perspective{
    UID: res.AssignedUids["perspective"],
    UIDString: uidString(res.AssignedUids["perspective"]),
    Statement: statement,
  }
  return perspective, nil
}
