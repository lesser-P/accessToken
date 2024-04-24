package auth

import (
	"accessToken/global"
	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	TOKEN    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjEyMywiaXNGcmVlemUiOmZhbHNlLCJTdGFuZGFyZENsYWltIjp7ImV4cCI6MTcxMjc0MjU4NX19.O1YEArK09y0epl16HUX-8hFy8FkNxYT15qhxstkn5VQ"
	USERID   = 123
	ISFREEZE = false
)

var (
	jwtInfo = global.GAL_Config.Jwt
	ctx     = getCtx()
)

func getCtx() *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	return ctx
}

func TestCreateToken(t *testing.T) {
	convey.Convey("准备 ctx，userId，isFreeze，issuer，key", t, func() {
		issuer := jwtInfo.Issuer
		key := jwtInfo.Key
		convey.Convey("调用CreateToken函数", func() {
			result, err := CreateToken(ctx, USERID, ISFREEZE, issuer, key)
			convey.Convey("结果必须是err为nil，result符合token规范", func() {
				convey.So(err, convey.ShouldEqual, nil)
				flage := len(result) > 0
				convey.So(flage, convey.ShouldBeTrue)
			})
		})
	})
}

func TestParseToken(t *testing.T) {
	convey.Convey("准备 ctx，token，key，userId", t, func() {
		key := jwtInfo.Key
		issuer := jwtInfo.Issuer
		token, createErr := CreateToken(ctx, USERID, ISFREEZE, issuer, key)
		convey.So(createErr, convey.ShouldBeEmpty)

		convey.Convey("调用ParseToken函数", func() {
			claim, err := ParseToken(ctx, token, key)
			userId := claim.Userid
			isFreeze := claim.IsFreeze
			convey.Convey("结果必须是userId和isFreeze为最初定义的值，err为nil", func() {
				convey.So(err, convey.ShouldBeEmpty)
				convey.So(userId, convey.ShouldEqual, USERID)
				convey.So(isFreeze, convey.ShouldEqual, ISFREEZE)
			})
		})
	})
}

func TestRefreshTokenStatus(t *testing.T) {
	convey.Convey("准备ctx，Token，key，创建一个token", t, func() {
		key := jwtInfo.Key
		issuer := jwtInfo.Issuer
		oldToken, createErr := CreateToken(ctx, USERID, ISFREEZE, issuer, key)
		time.Sleep(time.Second * 1)
		convey.So(createErr, convey.ShouldBeEmpty)

		convey.Convey("调用RefreshTokenStatus函数", func() {
			token, reFreshErr := RefreshTokenStatus(ctx, oldToken, key)
			convey.So(reFreshErr, convey.ShouldBeEmpty)

			convey.Convey("调用ParseToken函数", func() {
				newClaim, newParseErr := ParseToken(ctx, token, key)
				oldClaim, oldParseErr := ParseToken(ctx, oldToken, key)
				convey.So(newParseErr, convey.ShouldBeEmpty)
				convey.So(oldParseErr, convey.ShouldBeEmpty)

				newExpires := newClaim.StandardClaim.ExpiresAt
				oldExpires := oldClaim.StandardClaim.ExpiresAt
				convey.Convey("结果必须是reFreshErr,newParseErr，oldParseErr为nil,newExpires大于oldExpires", func() {
					isAfter := newExpires.Time.After(oldExpires.Time)
					convey.So(isAfter, convey.ShouldBeTrue)
				})
			})
		})
	})
}
