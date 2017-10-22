package db

import (
  "log"
)

// Comment TODO
type Comment struct {
  Text string `json:"text"`
  UID uint64 `json:"_uid_"`
  UIDString string `json:"uidstring"`
}

type newCommentQuery struct {
  Text string
  UserUID string
  SubjectUIDString string
}

func newComment(user *User, subjectUIDString string, text string) (*Comment, error) {
  data := newCommentQuery{
    Text: text,
    UserUID: user.UIDString,
    SubjectUIDString: subjectUIDString,
  }
  result, err := Query("new-comment", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  comment := &Comment{
    Text: text,
    UID: result.AssignedUids["comment"],
    UIDString: uidString(result.AssignedUids["comment"]),
  }
  return comment, nil
}

// NewComment will create a new comment by the given user on the given discussion. 
func (d *Discussion) NewComment(user *User, text string) (*Comment, error){
  return newComment(user, d.UIDString, text)
}

// NewReply will create a new reply on the given comment by the given user.
func (c *Comment) NewReply(user *User, text string) (*Comment, error){
  return newComment(user, c.UIDString, text)
}
