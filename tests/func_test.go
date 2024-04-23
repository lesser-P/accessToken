package tests

import (
	"accessToken_go_zero/internal/types"
	"context"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	convey.Convey("测试CreateToken方法", t, func() {
		token := types.Token{
			Key:     "ye",
			Issuer:  "0xYe",
			Timeout: 8,
		}
		newToken := types.CreateToken("ye", "0xYe", 8)
		convey.So(&token, convey.ShouldEqual, newToken)
	})
}
func TestGenerateToken(t *testing.T) {
	convey.Convey("测试GenerateToken", t, func() {
		jwt := types.CreateToken("ye", "0xYe", 8)
		ctx := context.Background()
		token, err := jwt.GenerateToken(ctx, "123", false)
		convey.Convey("验证生成token正确性", func() {
			convey.So(len(token), convey.ShouldBeGreaterThan, 0)
			convey.So(err, convey.ShouldBeNil)
			flag := jwt.VailToken(ctx, token)
			convey.So(flag, convey.ShouldBeTrue)
		})
	})
}
func TestFuncResetToken(t *testing.T) {
	convey.Convey("测试ResetToken", t, func() {
		ctx := context.Background()
		// 过期Token
		const OLDTOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyaWQiOiIzMyIsIklzRnJlZXplIjpmYWxzZSwiU3RhbmRhcmRDbGFpbSI6eyJpc3MiOiIweFllIiwiZXhwIjoxNzEzNzg0ODM0fX0.1uujeB3-zRZpHIhay0Z-W0QxLBRv4HtrMJyDkYBw1aM"
		tokenDetail := types.CreateToken("ye", "0xYe", 8)
		token, err := tokenDetail.ResetToken(ctx, OLDTOKEN)
		convey.Convey("测试过期token", func() {
			convey.So(token, convey.ShouldEqual, "")
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("测试重置有效期内Token", func() {
			generateToken, err := tokenDetail.GenerateToken(ctx, "0xYe", false)
			convey.So(len(generateToken), convey.ShouldBeGreaterThan, 0)
			convey.So(err, convey.ShouldBeNil)
			time.Sleep(time.Second * 2)
			resetToken, err := tokenDetail.ResetToken(ctx, generateToken)
			convey.So(len(resetToken), convey.ShouldBeGreaterThan, 0)
			convey.So(resetToken, convey.ShouldNotEqual, generateToken)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
func TestFuncVerifyToken(t *testing.T) {
	convey.Convey("测试VerifyToken方法", t, func() {
		jwt := types.CreateToken("0xYe", "0xYe", 8)
		ctx := context.Background()
		token, err := jwt.GenerateToken(ctx, "123", false)
		convey.Convey("判断token是否生成", func() {
			convey.So(len(token), convey.ShouldBeGreaterThan, 0)
			convey.So(err, convey.ShouldBeNil)
		})
		flag := jwt.VailToken(ctx, token)
		convey.Convey("判断刚生成的token是否可验证", func() {
			convey.So(flag, convey.ShouldBeTrue)
		})
		// 非法token
		const OLDTOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyaWQiOiJ0ZXN0IiwiSXNGcmVlemUiOmZhbHNlLCJTdGFuZGFyZENsYWltIjp7ImlzcyI6IjB4WWUiLCJleHAiOjE3MTM3ODQwMzV9fQ.ttN2PjBlACMWtzJqKZH9gozJ7Cfq6v9C28EwCvjilSk"
		flag = jwt.VailToken(ctx, OLDTOKEN)
		convey.Convey("判断非法的token是否可验证", func() {
			convey.So(flag, convey.ShouldBeFalse)
		})
	})
}
