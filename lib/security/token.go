package security

import (
  "github.com/dgrijalva/jwt-go"
  "time"
  "errors"
  "log"

  "diserve.didactia.org/lib/env"
  "diserve.didactia.org/lib/db"
)

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

// VerifyToken will given a JWT token, and a loginType as a string, returns
// true if the logintype matches and the token has not expired, false otherwise.
func VerifyToken(tokenString string, user *db.User) bool {
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
