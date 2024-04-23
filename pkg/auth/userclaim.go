package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

func NewEmptyClaim() *UserClaim {
	return &UserClaim{}
}

type UserClaim struct {
	Userid        string
	IsFreeze      bool
	StandardClaim jwt.RegisteredClaims
}

func (u UserClaim) GetExpirationTime() (*jwt.NumericDate, error) {
	//TODO implement me
	return u.StandardClaim.ExpiresAt, nil
}

func (u UserClaim) GetIssuedAt() (*jwt.NumericDate, error) {
	//TODO implement me
	return u.StandardClaim.IssuedAt, nil
}

func (u UserClaim) GetNotBefore() (*jwt.NumericDate, error) {
	//TODO implement me
	return u.StandardClaim.NotBefore, nil
}

func (u UserClaim) GetIssuer() (string, error) {
	//TODO implement me
	return u.StandardClaim.Issuer, nil
}

func (u UserClaim) GetSubject() (string, error) {
	//TODO implement me
	return u.StandardClaim.Subject, nil
}

func (u UserClaim) GetAudience() (jwt.ClaimStrings, error) {
	//TODO implement me
	return u.StandardClaim.Audience, nil
}
