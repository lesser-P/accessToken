package system

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	Userid        int  `json:"userid"`
	IsFreeze      bool `json:"isFreeze"`
	StandardClaim jwt.RegisteredClaims
}

func NewEmptyClaim() *UserClaim {
	return &UserClaim{}
}

func (u UserClaim) GetExpirationTime() (*jwt.NumericDate, error) {
	return u.StandardClaim.ExpiresAt, nil
}

func (u UserClaim) GetIssuedAt() (*jwt.NumericDate, error) {
	return u.StandardClaim.IssuedAt, nil
}

func (u UserClaim) GetNotBefore() (*jwt.NumericDate, error) {
	return u.StandardClaim.NotBefore, nil
}

func (u UserClaim) GetIssuer() (string, error) {
	return u.StandardClaim.Issuer, nil
}

func (u UserClaim) GetSubject() (string, error) {
	return u.StandardClaim.Subject, nil
}

func (u UserClaim) GetAudience() (jwt.ClaimStrings, error) {
	return u.StandardClaim.Audience, nil
}
