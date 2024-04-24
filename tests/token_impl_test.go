package types

import (
	"accessToken_go_zero/mocks"
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenToken(t *testing.T) {
	//ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mapper := mocks.NewMockTokenDao(ctrl)
	//sv := svc.ServiceContext{
	//	Mapper: mapper,
	//}

	convey.Convey("生成Token", t, func() {
		//token.NewGenTokenLogic(ctx, &sv)
		mapper.EXPECT().SaveOrUpdate(gomock.Any()).Return(nil).AnyTimes()
		//result, err := service.GenToken(&req)
		//convey.Convey("判断token不为空，错误为空", func() {
		//	convey.So(len(result), convey.ShouldBeGreaterThan, 0)
		//	convey.So(err, convey.ShouldBeNil)
		//})
	})
}

//func TestResetToken(t *testing.T) {
//	ctx := context.Background()
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	mapper := mocks.NewMockTokenDao(ctrl)
//	sv := svc.ServiceContext{
//		Mapper: mapper,
//		Token:  NewTokenFactory("0xYe", "0xYe", 8),
//	}
//	req := GenTokenReq{
//		UserId: "ye",
//	}
//	convey.Convey("生成Token", t, func() {
//		genService := token.NewGenTokenLogic(ctx, &sv)
//		mapper.EXPECT().SaveOrUpdate(gomock.Any()).Return(nil).AnyTimes()
//		result, err := genService.GenToken(&req)
//		convey.Convey("判断token不为空，错误为空", func() {
//			convey.So(len(result), convey.ShouldBeGreaterThan, 0)
//			convey.So(err, convey.ShouldBeNil)
//		})
//		convey.Convey("重置Token", func() {
//			tokenReq := TokenReq{
//				Token:  result,
//				UserId: "ye",
//			}
//			data := model.TokenDetail{}
//			resetService := token.NewResetTokenLogic(ctx, &sv)
//			mapper.EXPECT().SelectByToken(gomock.Any()).Return(&data, nil).AnyTimes()
//			resetToken, ResetErr := resetService.ResetToken(&tokenReq)
//
//			convey.So(len(resetToken.Data), convey.ShouldBeGreaterThan, 0)
//			convey.So(ResetErr, convey.ShouldBeNil)
//
//		})
//	})
//}
//
//func TestFreezeToken(t *testing.T) {
//	convey.Convey("冻结Token测试", t, func() {
//		ctx := context.Background()
//		ctrl := gomock.NewController(t)
//		defer ctrl.Finish()
//		mapper := mocks.NewMockTokenDao(ctrl)
//		sv := svc.ServiceContext{
//			Mapper: mapper,
//			Token:  NewTokenFactory("0xYe", "0xYe", 8),
//		}
//		mapper.EXPECT().FreezeToken(gomock.Any()).Return(nil)
//		service := token.NewFreezeTokenLogic(ctx, &sv)
//		tokenReq := TokenReq{
//			Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyaWQiOiJ5ZSIsIklzRnJlZXplIjpmYWxzZSwiU3RhbmRhcmRDbGFpbSI6eyJpc3MiOiIweFllIiwiZXhwIjoxNzEzODY5NjM3fX0.3ExLGpmvnOoz_S2wWNAhWhrvn7K-51REBRSa5XG_fXE",
//			UserId: "ye",
//		}
//		resp, err := service.FreezeToken(&tokenReq)
//		convey.Convey("判断结果", func() {
//			convey.So(err, convey.ShouldBeNil)
//			convey.So(resp.Code, convey.ShouldEqual, 200)
//		})
//	})
//}
//
//func TestVerifyToken(t *testing.T) {
//	convey.Convey("测试Token有效性", t, func() {
//		sv := svc.ServiceContext{
//			Token: NewTokenFactory("0xYe", "0xYe", 8),
//		}
//		jwt := NewTokenFactory("0xYe", "0xYe", 8)
//		ctx := context.Background()
//
//		result, err := jwt.GenerateToken(ctx, "123", false)
//		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
//		convey.So(err, convey.ShouldBeNil)
//		convey.Convey("调用验证函数", func() {
//			req := TokenReq{
//				Token: result,
//			}
//			service := token.NewVerifyTokenLogic(ctx, &sv)
//			resp, err := service.VerifyToken(&req)
//			convey.So(resp.Code, convey.ShouldEqual, 200)
//			convey.So(err, convey.ShouldBeNil)
//		})
//	})
//}
