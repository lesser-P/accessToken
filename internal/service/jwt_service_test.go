package service

import (
	"accessToken/global"
	mock_dao "accessToken/mock"
	"accessToken/model/system"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)

const TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjEyMywiaXNGcmVlemUiOmZhbHNlLCJTdGFuZGFyZENsYWltIjp7ImV4cCI6MTcxMjgxMzMxMn19.AToi8t2_Vq-xz7f81vt_UtSNFefdsGfTR_ITJ5cQHZ0"

func TestGenerateToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	//expectedErr := errors.New("保存信息异常：err")

	mockTokenDao := mock_dao.NewMockTokenDao(mockCtrl)
	detail := &system.TokenDetails{
		Token:  TOKEN,
		UserId: 123,
	}
	// 配置mock对象的期望行为
	mockTokenDao.EXPECT().GetRecordById(ctx, gomock.Eq(123)).Return(detail, nil).Times(1)
	//mockTokenDao.EXPECT().
	//	CreateTokenRecord(ctx, gomock.Eq(TOKEN), gomock.Eq(123)).
	//	Return(expectedErr).Times(1)
	service := NewJwtService(ctx, global.GAL_DB)
	service.dao = mockTokenDao
	if _, err := service.GenerateToken(ctx, 123); err != nil {
		t.Errorf("faild to generateToken %v", err)
	}
}
