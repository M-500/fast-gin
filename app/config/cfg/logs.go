package cfg

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 11:58
//

type LogCfg struct {
	Level         string `mapstructure:"level"`
	Path          string `mapstructure:"path"`
	Filename      string `mapstructure:"filename"`
	MaxSize       int    `mapstructure:"max_size"`
	MaxAge        int    `mapstructure:"max_age"`
	MaxBackups    int    `mapstructure:"max_backups"`
	Compress      bool   `mapstructure:"compress"`
	StacktraceKey string `mapstructure:"stacktrace_key"`
	Format        string `mapstructure:"format"`
}
