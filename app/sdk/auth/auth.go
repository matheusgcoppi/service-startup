// Package auth provides authentication and authorization support.
// Authentication: You are who you say you are.
// Authorization: You have permission to do what you are requesting to do.
package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/matheusgcoppi/service/foundation/logger"
)

// ErrForbidden is returned when an auth issue is identified
var ErrForbidden = errors.New("attempted action is not allowed")

// Claims represent the authorization claims transmitted via a jwt.
type Claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

// HasRole checks if the specified role exists.
func (c Claims) HasRole(r string) bool {
	for _, role := range c.Roles {
		if role == r {
			return true
		}
	}
	return false
}

// KeyLookup declares a method set of behavior for looking up
// private and public keys for JWT use. The return could be a
// PEM encoded string or a JWT-based key.
type KeyLookup interface {
	PrivateKey(kid string) (key string, err error)
	PublicKey(kid string) (key string, err error)
}

// Config represents information required to initialize auth.
type Config struct {
	Log       *logger.Logger
	KeyLookup KeyLookup
	Issuer    string
}

// Auth is used to authenticate clients. It can generate a token for a
// set of user claims and recreate the claims by parsing the token.
type Auth struct {
	keyLookup KeyLookup
	method    jwt.SigningMethod
	parser    *jwt.Parser
	issued    string
}

//// New creates an Auth to support authentication/authorization.
//func New(cfg Config) (*Auth, error) {
//
//	If a database connection is not provided, we won't perform the user enabled check.
//	var userBus *userbus.Core
//	if cfg.DB != nil {
//		userBus
//	}
//
//	a := Auth{
//		keyLookup: cfg.KeyLookup,
//	}
//}
