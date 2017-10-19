package db

import (
  "errors"
  "fmt"
)

// ErrEmptyStatementSlice self explanatory
var ErrEmptyStatementSlice = errors.New("cannot generate statements from empty list")

// ErrIncompleteUIDMap self explanatory
var ErrIncompleteUIDMap = errors.New("cannot find label in assigned uids")

// Statement TODO
type Statement struct {
  UID uint64 `json:"_uid_"`
  UIDString string `json:"uidstring"`
  Text string `json:"text"`
  Next *Statement `json:"next"`
}

type idAndText struct {
  ID string
  Text string
}

func identifyTexts(texts []string) ([]idAndText) {
  var output []idAndText
  output = make([]idAndText, len(texts))
  for i, text := range texts {
    output[i].ID = fmt.Sprintf("s%d", i)
    output[i].Text = text
  }
  // used for checking for last in template.
  output[len(output)-1].ID = ""
  return output
}


func statementSerializer(iAndT []idAndText, aUids map[string]uint64) (*Statement, error) {
  // could be recursive instead, but it might have issues with long statements.
  if len(iAndT) == 0 {
    return nil, ErrEmptyStatementSlice
  }
  firstUID, ok := aUids[iAndT[0].ID]
  if !ok {
    return nil, ErrIncompleteUIDMap
  }
  stmt := Statement{
    Text: iAndT[0].Text,
    UID: firstUID,
    UIDString: uidString(firstUID),
  }
  first := stmt
  for _, val := range iAndT[1:len(iAndT)-1] {
    uid, ok := aUids[val.ID]
    if !ok {
      return nil, ErrIncompleteUIDMap
    }
    next := Statement{
      Text: val.Text,
      UID: uid,
      UIDString: uidString(aUids[val.ID]),
    }
    stmt.Next = &next
    stmt = next
  }
  if len(iAndT) > 1 {
    uid, ok := aUids["end"]
    if !ok {
      return nil, ErrIncompleteUIDMap
    }
    last := Statement{
      Text: iAndT[len(iAndT)].Text,
      UID: uid,
      UIDString: uidString(uid),
    }
    stmt.Next = &last
  }
  return &first, nil
}
