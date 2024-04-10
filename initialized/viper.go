package initialized

import (
	"accessToken/global"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("viper ReadInConfig err:" + err.Error())
	}
	err = v.Unmarshal(&global.GAL_Config)
	if err != nil {
		panic("viper Unmarshal err:" + err.Error())
	}
	return v
}
