package logic

import (
	"accessToken_grpc/internal/domain"
	"accessToken_grpc/internal/model"
	"accessToken_grpc/internal/svc"
	"accessToken_grpc/mocks"
	"accessToken_grpc/pb/token"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFreezeTokenLogic(t *testing.T) {
	convey.Convey("冻结Token测试", t, func() {
		ctx := context.Background()
		ctl := gomock.NewController(t)

		defer ctl.Finish()
		mapper := mocks.NewMockTokenDao(ctl)
		sv := svc.ServiceContext{Mapper: mapper}
		service := NewFreezeTokenLogic(ctx, &sv)
		mapper.EXPECT().SelectByToken(gomock.Any(), gomock.Any()).Return(&model.TokenDetail{IsFreeze: false}, nil)
		mapper.EXPECT().SaveOrUpdate(gomock.Any(), gomock.Any()).Return(nil)
		result, err := service.FreezeToken(&token.TokenReq{})
		convey.Convey("判断返回数据是否正常", func() {
			convey.So(result, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestGenTokenLogic_GenToken(t *testing.T) {
	convey.Convey("生成Token测试", t, func() {
		ctx := context.Background()
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		mapper := mocks.NewMockTokenDao(ctl)
		sv := svc.ServiceContext{Mapper: mapper}
		service := NewGenTokenLogic(ctx, &sv)
		mapper.EXPECT().SaveOrUpdate(gomock.Any(), gomock.Any()).Return(nil)
		jwt, err := service.GenToken(&token.GenReq{
			UserId: "berg",
		})
		convey.Convey("判断结果", func() {
			convey.So(jwt, convey.ShouldNotBeEmpty)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestRestTokenLogic_ResetToken(t *testing.T) {
	convey.Convey("重置Token测试", t, func() {
		ctx := context.Background()
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		mapper := mocks.NewMockTokenDao(ctl)
		sv := svc.ServiceContext{
			Mapper: mapper,
			Token:  *domain.NewTokenFactory("0xYe", "0xYe", 8),
		}
		service := NewResetTokenLogic(ctx, &sv)

		mapper.EXPECT().SelectByToken(gomock.Any(), gomock.Any()).Return(&model.TokenDetail{}, nil).AnyTimes()
		mapper.EXPECT().SaveOrUpdate(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
		req := token.TokenReq{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyaWQiOiJxcTEiLCJJc0ZyZWV6ZSI6ZmFsc2UsIlN0YW5kYXJkQ2xhaW0iOnsiaXNzIjoiMHhZZSIsImV4cCI6MTcxMzk0Njk1NX19.o1EjWOS8LZbarv1ZGQmzdTBCKYDEwF07wnLA1leyaas",
		}

		resetToken, err := service.ResetToken(&req)
		convey.Convey("判断重置是否成功", func() {
			convey.So(resetToken, convey.ShouldNotEqual, req.Token)
			convey.So(resetToken, convey.ShouldNotBeEmpty)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestVerifyTokenLogic_VerifyToken(t *testing.T) {
	convey.Convey("验证Token测试", t, func() {
		factory := domain.NewTokenFactory("0xYe", "0xYe", 8)
		ctx := context.Background()
		jwt, err := factory.GenerateToken(ctx, "re", false)
		convey.Convey("判断token是否生成", func() {
			convey.So(jwt, convey.ShouldNotBeEmpty)
			convey.So(err, convey.ShouldBeNil)
		})
		err = factory.VailToken(ctx, jwt)
		convey.Convey("判断token是否验证成功", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
