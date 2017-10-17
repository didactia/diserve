package db

import (
  "github.com/dgraph-io/dgraph/client"
  "strconv"
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
func (dbc *DatabaseClient) LoginUser(name string, password string) (*User, error) {
  data := loginQuery{
    Name: name,
    Password: password,
  }
  res, err := dbc.Query("login", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  var r loginResponse
  err = client.Unmarshal(res.N, &r)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseUnmarshalling
  }
  if r.Login.UID == 0 {
    return nil, ErrUsernameNotFound
  }
  if !r.Login.Password.Check {
    return nil, ErrPasswordIncorrect
  }
  user := &User{
    Name: r.Login.Name,
    UID: r.Login.UID,
    UIDString: strconv.FormatUint(r.Login.UID, 16),
  }
  return user, nil
}

// RegisterUser will save a user with the given name and password, returns the
// User struct, error will be nil if name is unique.
func (dbc *DatabaseClient) RegisterUser(name string, password string) (*User, error){
  _, err := dbc.GetUser(name)
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
  res, err := dbc.Query("register", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  user := &User{
    Name: name,
    UID: res.AssignedUids["user"],
    UIDString: strconv.FormatUint(res.AssignedUids["user"], 16),
  }
  return user, nil
}

// GetUser will given a user name return a user struct.
// error is nil if user exists.
func (dbc *DatabaseClient) GetUser(name string) (*User, error) {
  data := userQuery{
    Name: name,
  }
  result, err := dbc.Query("get-user", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  var r userResponse
  err = client.Unmarshal(result.N, &r)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseUnmarshalling
  }
  // User doesn't exist.
  if r.User.UID == 0 {
    return nil, ErrUsernameNotFound
  }
  user := &User{
    Name: r.User.Name,
    UID: r.User.UID,
    UIDString: strconv.FormatUint(r.User.UID, 16),
  }
  return user, nil
}
