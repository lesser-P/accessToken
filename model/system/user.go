package system

import (
	"accessToken/global"
	"accessToken/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId int64 `json:"userId"`
	Ctx    *gin.Context
	BaseModel
}

func CreateUserFactoryWithCtx(ctx *gin.Context) *User {
	return &User{
		Ctx: ctx,
		BaseModel: BaseModel{
			DB: global.GAL_DB,
		},
	}
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) AddUser(user *User) error {
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		factory := utils.CreateRandFactory(uint64(time.Now().UnixNano()), global.RAND_START, global.RAND_END)
		user.UserId = factory.GetRandInt()
		tx.Create(&user)
		if err := tx.Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
