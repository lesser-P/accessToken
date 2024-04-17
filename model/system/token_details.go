package system

type TokenDetails struct {
	UserId   int    `json:"userId" gorm:"column:user_id"`
	Token    string `json:"token" gorm:"column:token"`
	IsFreeze bool   `json:"isFreeze" gorm:"column:is_freeze"`
}

func (t TokenDetails) TableName() string {
	return "token_details"
}
