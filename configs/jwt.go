package configs

type Jwt struct {
	Key     string `json:"key" mapstructure:"key"`
	Expires string `json:"expires" mapstructure:"expires"`
	Issuer  string `json:"issuer" mapstructure:"issuer"`
}
