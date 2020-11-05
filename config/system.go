package config

type System struct {
	Port   string `mapstructure:"port" json:"port" yaml:"port"`
	DbType string `mapstructure:"dbType" json:"dbType" yaml:"dbType"`
}
