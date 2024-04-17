package auth

import (
	"accessToken/model/system"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

const Expire = 7200

func CreateToken(ctx *gin.Context, userId int, isFreeze bool, issuer string, key string) (string, error) {
	claims := system.UserClaim{
		Userid:   userId,
		IsFreeze: isFreeze,
		StandardClaim: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * Expire)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println("创建token失败：", err.Error())
		return "", err
	}
	return str, nil
}

func ParseToken(ctx *gin.Context, str string, key string) (*system.UserClaim, error) {
	claim := system.NewEmptyClaim()
	token, err := jwt.ParseWithClaims(str, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if claim.IsFreeze {
		return nil, errors.New("token已冻结，无法访问")
	}
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("异常token，无法访问")
	}
	return claim, nil
}

func ParseReturnToken(ctx *gin.Context, str string, key string) (*jwt.Token, error) {
	claim := system.NewEmptyClaim()
	return jwt.ParseWithClaims(str, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
}

func RefreshTokenStatus(ctx *gin.Context, str string, key string) (string, error) {
	claim := system.NewEmptyClaim()
	_, err := jwt.ParseWithClaims(str, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	claim.StandardClaim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * Expire))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	newToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return newToken, err
}
