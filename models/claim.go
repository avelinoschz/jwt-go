package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim contains the JWT payload information
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
