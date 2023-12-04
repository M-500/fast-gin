package cfg

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 11:57
//

type JwtCfg struct {
	SecretKey     string `mapstructure:"secret_key"`
	JwtHeaderKey  string `mapstructure:"jwt_header_key"`
	JwtPrefix     string `mapstructure:"jwt_prefix"`
	JwtExpireHour int64  `mapstructure:"jwt_expire_hour"`
}
