package config

type Mysql struct {
	HostName string `mapstructure:"hostName" json:"hostName" yaml:"hostName"`
	HostPort string `mapstructure:"hostPort" json:"hostPort" yaml:"hostPort"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
