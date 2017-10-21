package security

import (
  "github.com/dgrijalva/jwt-go"
  "time"
  "log"
  "fmt"
  "errors"

  "diserve.didactia.org/lib/util"
  "diserve.didactia.org/lib/env"
  "diserve.didactia.org/lib/db"
)

// ErrMissingUIDString occurs when given a struct without an UIDString.
var ErrMissingUIDString = errors.New("missing uid string")

// ErrUnexpectedSigningMethod is given if there is a signing method mismatch
// in the token header.
var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

// ErrClaimsSigning occurs in jwt signing, data safe.
var ErrClaimsSigning = errors.New("error on claims signing")

// LoginClaims are jwt standard claims, with LoginType (ltp in json)
type LoginClaims struct {
  jwt.StandardClaims
}

// LoginToken given a user, returns a JWT string, and its expiration.
func LoginToken(user *db.User) (string, int64, error) {
  now := time.Now().Unix()
  expiration := now + env.Vars.USERTOKENDURATION
  claims := LoginClaims{
    jwt.StandardClaims{
      IssuedAt: now,
      ExpiresAt: expiration,
      Subject: user.UIDString,
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString(env.Vars.HMACSECRET)
  if err != nil {
    log.Print(err)
    return "", 0, ErrClaimsSigning
  }
  return tokenString, expiration, nil
}


// VerifyLoginToken will given a JWT token, and a loginType as a string, returns
// true if the logintype matches and the token has not expired, false otherwise.
func VerifyLoginToken(tokenString string, user *db.User) bool {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, ErrUnexpectedSigningMethod
    }
    return env.Vars.HMACSECRET, nil
  })
  if err != nil {
    log.Print(err)
    return false
  }
  claims, ok := token.Claims.(jwt.MapClaims);
  return ok && token.Valid && claims["sub"] == user.UIDString
}

type structClaims struct {
  UIDString string
  StructType string
  jwt.StandardClaims
}

// StructUIDToken will given a struct with an UIDString, return a JWT certifying that the
// UID corresponds to that struct type.
func StructUIDToken(data interface{}) (string, error) {
  UIDString, err := util.GetStringField(data, "UIDString")
  if err != nil {
    return "", err
  }
  if UIDString == "" {
    return "", ErrMissingUIDString
  }
  claims := structClaims{
    UIDString: UIDString,
    StructType: util.TypeString(data),
  }
  fmt.Println(claims.UIDString)
  fmt.Println(claims.StructType)
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString(env.Vars.HMACSECRET)
  if err != nil {
    log.Print(err)
    return "", ErrClaimsSigning
  }
  return tokenString, nil
}

// VerifyStructUIDToken will given a JWT token, and a struct name, and an
// UIDString, verify that the JWT token corresponds to those.
func VerifyStructUIDToken(tokenString string, typeString string, UIDString string) bool {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, ErrUnexpectedSigningMethod
    }
    return env.Vars.HMACSECRET, nil
  })
  if err != nil {
    log.Print(err)
    return false
  }
  claims, ok := token.Claims.(jwt.MapClaims);
  return ok && claims["StructType"] == typeString && claims["UIDString"] == UIDString
}
