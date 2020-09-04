// grpc 服务入口

package main

import (
	"go-scaffold/internal/grpc"
	"go-scaffold/pkg/conf"
	"go-scaffold/pkg/log"
	"go-scaffold/pkg/xgorm"
	"go-scaffold/pkg/xredis"
)

func main() {
	// fmt.Printf("processID: %d\n", os.Getppid())
	conf.Init()                                             // 加载配置文件
	log.RawConfig("app.log", conf.GetHandler()).Init()      // 加载日志
	defer log.Sync()                                        // 日志落盘
	xgorm.RawConfig("app.mysql", conf.GetHandler()).Init()  // 加载gorm
	xredis.RawConfig("app.redis", conf.GetHandler()).Init() // 加载redis
	defer xredis.Cli.Close()
	grpc.Serve()
}
