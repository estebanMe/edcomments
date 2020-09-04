package models

import "github.com/dgrijalva/jwt-go"

//Claim User Token
type Claim struct {
	User `json:"user "`
	jwt.StandardClaims
}
