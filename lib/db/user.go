package db

import (
  "errors"
  "log"
)

// ErrUsernameNotFound is given when queried username is not found
var ErrUsernameNotFound = errors.New("username not found")

// ErrPasswordIncorrect is given when password is incorrect
var ErrPasswordIncorrect = errors.New("password is incorrect")

// ErrUsernameAlreadyExists is given on registration if user with username
// already exists
var ErrUsernameAlreadyExists = errors.New("username already exists")


// User TODO
type User struct {
  Name string `json:"name"`
  UID uint64 `json:"_uid_"`
  UIDString string
}

type password struct {
  Check bool `json:"checkpwd"`
}

type login struct {
  Name string   `json:"name"`
  UID uint64    `json:"_uid_"`
  Password *password `json:"password"`
}

type loginQuery struct {
  Name string
  Password string
}

type loginResponse struct {
  Login *login `json:"login"`
}

type userQuery struct {
  Name string
}

type userResponse struct {
  User *User `json:"user"`
}

type uid struct {
  Value uint64 `json:"user"`
}

type registerResponse struct {
  UID *uid `json:"AssignedUids"`
}

// LoginUser will try to login with the name and password, returns the User
// struct. Errors returned are stripped for data.
func LoginUser(name string, password string) (*User, error) {
  data := loginQuery{
    Name: name,
    Password: password,
  }
  var res loginResponse
  err := QueryAndUnmarshal("login", data, &res)
  if err != nil {
    return nil, err
  }
  if res.Login.UID == 0 {
    return nil, ErrUsernameNotFound
  }
  if !res.Login.Password.Check {
    return nil, ErrPasswordIncorrect
  }
  user := &User{
    Name: res.Login.Name,
    UID: res.Login.UID,
    UIDString: uidString(res.Login.UID),
  }
  return user, nil
}

// NewUser will save a user with the given name and password, returns the
// User struct, error will be nil if name is unique.
func NewUser(name string, password string) (*User, error){
  _, err := GetUser(name)
  if err == nil {
    return nil, ErrUsernameAlreadyExists
  }
  if err != ErrUsernameNotFound {
    return nil, err
  }
  data := loginQuery{
    Name: name,
    Password: password,
  }
  res, err := Query("register", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  user := &User{
    Name: name,
    UID: res.AssignedUids["user"],
    UIDString: uidString(res.AssignedUids["user"]),
  }
  return user, nil
}

// GetUser will given a user name return a user struct.
// error is nil if user exists.
func GetUser(name string) (*User, error) {
  data := userQuery{
    Name: name,
  }
  var res userResponse
  err := QueryAndUnmarshal("get-user", data, &res)
  if err != nil {
    return nil, err
  }
  // User doesn't exist.
  if res.User.UID == 0 {
    return nil, ErrUsernameNotFound
  }
  user := &User{
    Name: res.User.Name,
    UID: res.User.UID,
    UIDString: uidString(res.User.UID),
  }
  return user, nil
}
