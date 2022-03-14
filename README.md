## 架构

目录结构如下：

```shell
blog-service
├── configs		// 配置文件
├── docs		// 文档集合
├── global		// 全局变量
├── internal	// 内部模块
│   ├── dao				// 数据访问层Database Access Object
│   ├── middleware		// HTTP 中间件
│   ├── model			// 模型层，用于存放 model 对象
│   ├── routers			// 路由相关逻辑处理
│   └── service			// 项目核心业务逻辑
├── pkg			// 项目相关的模块包
├── storage		// 项目生成的临时文件
├── scripts		// 各类构建，安装，分析等操作的脚本
└── third_party	// 第三方的资源工具，例如 Swagger UI
```


## 使用包

* 路由：`gin`
* 配置管理：`viper`
* 数据库操作：`gorm`
* 日志：`zap`，`lumberjack`
* 接口文档生成：`swaggo`
* 接口参数校验：`validator`
* 访问控制：`jwt-go`
* 邮件提醒：`gomail.v2`
* 流量控制：`ratelimit`

## 内部全局部件

### log

日志使用方法

如：

```go
global.Logger.Info("你好")
global.Logger.WarnF("%s is %s", "kk", "biss")
```

### swagger

更新接口文档后，终端执行`swag init`初始化接口文档

### errorCode

定义于`/pkg/errorCode`中，内部对错误码处理需与项目全局统一

### 类型转换

定义于`/pkg/convert`中，统一全局的类型转换
