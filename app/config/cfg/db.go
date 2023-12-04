package cfg

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 11:55

type MysqlCfg struct {
	Datasource string `mapstructure:"datasource"`
}

type RedisCfg struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Pass string `mapstructure:"pass"`
}
