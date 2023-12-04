package config

import "fast-gin/app/config/cfg"

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 11:55
var ConfigInstance *Config

type Config struct {
	Name  string       `mapstructure:"name"`
	Host  string       `mapstructure:"host"`
	Local string       `mapstructure:"local"`
	Port  int          `mapstructure:"port"`
	MySQl cfg.MysqlCfg `mapstructure:"mysql"`
	Redis cfg.RedisCfg `mapstructure:"redis"`
	Jwt   cfg.JwtCfg   `mapstructure:"jwt"`
}
