package configs

type System struct {
	Port         string `json:"port" mapstructure:"port"`
	RouterPrefix string `json:"routerPrefix" mapstructure:"router-prefix"`
}
