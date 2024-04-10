package auth

import (
	"accessToken/global"
	"fmt"
	"log"
	"testing"
)

var jwtInfo = global.GAL_Config.Jwt

func TestCreateToken(t *testing.T) {
	token, err := CreateToken(12, false, jwtInfo.Issuer, jwtInfo.Key)
	if err != nil {
		return
	}
	fmt.Print(token)
}

func TestParseToken(t *testing.T) {
	userClaim, err := ParseToken("token", jwtInfo.Key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	userid := userClaim.Userid
	freezed := userClaim.IsFreeze
	fmt.Println(userid)
	log.Println(freezed)
}

func TestRefreshTokenStatus(t *testing.T) {
	token, err := CreateToken(12, false, jwtInfo.Issuer, jwtInfo.Key)
	if err != nil {
		return
	}
	newToken, err := RefreshTokenStatus(token, jwtInfo.Key)
	if err != nil {
		log.Println("err:", err.Error())
		return
	}
	log.Println(newToken)
}
