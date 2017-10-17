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

type password struct {
  Check bool `dgraph:"checkpwd"`
}

type login struct {
  Name string   `dgraph:"name"`
  UID uint64    `dgraph:"_uid_"`
  Password *password `dgraph:"password"`
}

// User TODO
type User struct {
  Name string
  UID uint64
  UIDString string
}

type loginQueryData struct {
  Name string
  Password string
}

type loginResponse struct {
  Login *login `dgraph:"login"`
}
// LoginUser will try to login with the name and password, returns the User
// struct. Errors returned are stripped for data.
func (dbc *DatabaseClient) LoginUser(name string, password string) (*User, error) {
  data := loginQueryData{
    Name: name,
    Password: password,
  }
  result, err := dbc.Query("login", data)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseQuery
  }
  var l loginResponse
  err = client.Unmarshal(result.N, &l)
  if err != nil {
    log.Print(err)
    return nil, ErrResponseUnmarshalling
  }
  if l.Login.UID == 0 {
    return nil, ErrUsernameNotFound
  }
  if !l.Login.Password.Check {
    return nil, ErrPasswordIncorrect
  }
  user := &User{
    Name: l.Login.Name,
    UID: l.Login.UID,
    UIDString: strconv.FormatUint(l.Login.UID, 16),
  }
  return user, nil
}

// RegisterUser will save a user with the given name and password, returns the
// User struct, error will be nil if name is unique.
func RegisterUser(name string, password string) (*User, error){
  return nil, nil
}
