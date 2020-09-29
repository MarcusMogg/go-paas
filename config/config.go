package config

// Config 所有配置
type Config struct {
	AesKey     string `mapstructure:"asekey" json:"asekey" yaml:"asekey"`
	JWTKey     string `mapstructure:"jwtkey" json:"jwtkey" yaml:"jwtkey"`
	Mysql      Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Addr       int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DockerHost string `mapstructure:"dockerhost" json:"dockerhost" yaml:"dockerhost"`
	Portbase   int    `mapstructure:"port" json:"port" yaml:"port"`
}

// Mysql 连接配置
type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Parm     string `mapstructure:"parm" json:"parm" yaml:"parm"`
}
