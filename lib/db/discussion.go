package db

import (
  "log"
)

// Discussion TODO
type Discussion struct {
  Text string `json:"text"`
  DTitle string `json:"dtitle"`
  UID uint64 `json:"_uid_"`
  UIDString string `json:"uidstring"`
}

type newDiscussionQuery struct {
  Text string
  DTitle string
  UserUIDString string
  SubjectUIDString string
}

// NewDiscussion creates a new discussion for the subject corresponding to the
// uidString, this discussion will have the given title. title is not unique.
func NewDiscussion(user *User, dTitle string, text string, subjectUIDString string) (*Discussion, error) {
  data := newDiscussionQuery{
    Text: text,
    DTitle: dTitle,
    UserUIDString: user.UIDString,
    SubjectUIDString: subjectUIDString,
  }
  result, err := Query("new-discussion", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  discussion := &Discussion{
    DTitle: dTitle,
    UID: result.AssignedUids["discussion"],
    UIDString: uidString(result.AssignedUids["discussion"]),
  }
  return discussion, nil
}
