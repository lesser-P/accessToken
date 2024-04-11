package service

import (
	mock_dao "accessToken/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	expectedErr := errors.New("该用户ID已存在token")

	mockTokenDao := mock_dao.NewMockTokenDao(mockCtrl)

	// 配置mock对象的期望行为
	mockTokenDao.EXPECT().
		CreateTokenRecord("111", 123).
		Return(expectedErr).
		Times(1)
	service := NewJwtService()
	service.dao = mockTokenDao
	if _, err := service.GenerateToken(123); err != nil {
		t.Errorf("faild to generateToken %v", err)
	}
}
