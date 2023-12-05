//@Author: wulinlin
//@Description:
//@File:  main
//@Version: 1.0.0
//@Date: 2023/12/01 22:55

package main

import (
	"fast-gin/app/config"
	"fast-gin/app/svc"
	"fast-gin/pkg/utils"
	"fast-gin/routers"
	"flag"
	"fmt"
)

var configFile = flag.String("f", "dev.yaml", "the config file")

func main() {

	utils.DisplayFastGin()
	// 1. 初始化配置文件
	flag.Parse()
	serverCfg := config.MustLoadCfg(*configFile, "YML")
	fmt.Printf(serverCfg.Name)
	svc.SetUpServiceContext(serverCfg)

	r := routers.SetUpRouters()
	panic(r.Run(fmt.Sprintf("%s:%d", serverCfg.Host, serverCfg.Port)))
}
