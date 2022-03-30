## 架构

目录结构如下：

```shell
blog-service
├── configs     // 配置文件
├── DB          // 用于存放数据库文件
├── docs        // 文档集合
├── global      // 全局变量
├── internal    // 内部模块
│   ├── dao             // 数据访问层Database Access Object
│   ├── middleware      // HTTP 中间件
│   ├── model           // 模型层，用于存放 model 对象
│   ├── routers         // 路由相关逻辑处理
│   └── service         // 项目核心业务逻辑
├── pkg			// 项目相关的模块包
├── scripts		// 各类构建，安装，分析等操作的脚本
├── storage		// 项目生成的临时文件
├── third_party	// 第三方的资源工具，例如 Swagger UI
├── Dockerfile  // Docker的加载文件
├── Goose       // 编译文件
├── main.go     // 入口文件
└── README.md   // 项目介绍
```
## 项目介绍
该项目为一个论坛系统的后端系统，当前项目仍处于开发中，已将大概框架架构完成，在逐步补充API，开发重心将暂时偏移向前端（以及学习操作系统上）。

## 项目部署
目前项目在Docker上进行了部署，因此为了更好的进行这个项目的部署，最好了解：
* Go语言的编译运行和包导入
* Docker的基本使用
* MySQL和Redis的基本使用

当然，接下来的部署流程将较为详细的介绍整个部署的过程，因而如果你对上面的知识不大了解也可以将项目运行起来。
### 部署流程
#### Docker镜像导入
* `docker pull gwhanmur/goose:v1.3`
* `docker pull redis`
* `docker pull mysql`

#### Redis启动
```shell
docker run --name redis -p 6380:6379 redis
```
其中,
* `-–name`：容器名，此处命名为mysql
* `-p`：端口映射，此处映射 主机3307端口 到 容器的3306端口

#### MySQL启动
##### 数据库启动
终端进入项目的`DB`文件夹中
```shell
docker run -p 3307:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```
其中,
* `-–name`：容器名，此处命名为mysql
* `-e`：配置信息，此处配置mysql的root用户的登陆密码
* `-p`：端口映射，此处映射 主机3307端口 到 容器的3306端口
##### 数据表导入
将数据库文件导入容器
```shell
docker cp goose.sql goose:/root/
```
进入容器
```shell
docker exec -it mysql bash
```
导入数据库
```shell
mysql -uroot -p123456 < goose.sql
```
> 这里的root为你的数据库账号，123456为你的数据库密码

#### 服务运行
```shell
docker run --link mysql:mysql --link redis:redis --name goose -p 8000:8000 goose:v1.3
```
其中，
* `-–name`：容器名，此处命名为goose
* `-p`：端口映射，此处映射 主机8000端口 到 容器的8000端口
* `--link [container'name]:[inner'name]` 是将已有的container注册为服务内部的名称，使用方式如`mysql:3306`

#### 项目访问
项目启动于 `localhost:8000` , 通过`localhost:8000/swagger/index.html` 可以访问其api文档并在其中进行一些api的测试

## 项目相关
### 使用包
* 路由：`gin`
* 配置管理：`viper`
* 数据库操作：`gorm`
* 日志：`zap`，`lumberjack`
* 接口文档生成：`swaggo`
* 接口参数校验：`validator`
* 访问控制：`jwt-go`
* 邮件提醒：`gomail.v2`
* 流量控制：`ratelimit`

### 内部全局部件

#### log

日志使用方法

如：

```go
global.Logger.Info("你好")
global.Logger.WarnF("%s is %s", "kk", "biss")
```

#### swagger

更新接口文档后，终端执行`swag init`初始化接口文档

#### errorCode

定义于`/pkg/errorCode`中，内部对错误码处理需与项目全局统一

#### 类型转换

定义于`/pkg/convert`中，统一全局的类型转换
