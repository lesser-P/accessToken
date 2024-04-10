package configs

type Mysql struct {
	Path     string `json:"path" mapstructure:"path"`
	Port     string `json:"port" mapstructure:"port"`
	Dbname   string `json:"dbName" mapstructure:"db-name"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

func (m *Mysql) DNS() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname
}
