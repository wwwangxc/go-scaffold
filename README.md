# go-scaffold/

https://github.com/etcd-io/etcd/pull/11564/files

### 目录结构

```shell
.
├── cmd                   # 应用程序入口
│   ├── grpc              # grpc服务入口
│   ├── grpc_test
│   └── http              # http服务入口
├── config                # 配置文件
├── internal              # 业务代码
│   ├── constant          # 全局常数
│   ├── grpc              # grpc服务
│   │   ├── handler       # 服务实现
│   │   └── pb            # protobuf契约及自动生成的契约代码
│   ├── http              # http服务
│   │   ├── dto	          # 出入参实体
│   │   ├── handler       # 服务实现
│   │   ├── middleware    # 中间件
│   │   └── swagger       # swagger文档（自动生成）
│   ├── model             # 数据库实体
│   └── service           # 业务实现
├── pkg                   # 封装的第三方库及工具
│   ├── conf              # 配置文件
│   ├── etcd              # etcd解释器及注册机
│   ├── log               # 日志
│   ├── util              # 工具
│   ├── xgin              # gin
│   ├── xgorm             # gorm
│   ├── xgrpc             # grpc
│   ├── xmysql            # mysql driver
│   └── xredis            # redis
├── script                # 脚本文件
└── target                # 编译后生成的目标文件及配置
    ├── grpc              # grpc服务文件
    └── http              # http服务文件
```

