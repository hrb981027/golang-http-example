package config

type Redis struct {
	HostName string `mapstructure:"hostName" json:"hostName" yaml:"hostName"`
	HostPort string `mapstructure:"hostPort" json:"hostPort" yaml:"hostPort"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
