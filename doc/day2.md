这是我参与「第五届青训营」伴学笔记创作活动的第2天

# 1.课前准备
安装开发需要的软件: Docker; Postman; Git; Golang
其中 Postman 是一款浏览器插件。需要在chrome extensions里安装
# 2.重点内容
- GORM： 功能强大的ORM框架
- Hertz: 微服务 HTTP 框架，具有高易用性、高性能、高扩展性等特点
- Kitex: 微服务 RPC 框架，具有高性能、强可扩展的特点<br>
注释：
- ORM (Object Relational Mapping) 在关系型数据库和对象之间作一个映射
- RPC (Remote Produce Call) 远程过程调用,它是一种通过网络从远程计算机程序上请求服务,而不需要了解底层网络技术的协议
# 3.详细介绍
## gorm 的约定
使用 ID 作为主键；使用结构体的 **snake_cases** (复数) 作为表名；使用 struct 字段名的 snake_case 作为列名；使用CreatedAt，
UpdatedAt 作为创建时间<br>
gorm 通过驱动连接数据库，支持MySQL，SQLite。<br>
[DSN (Data Source Name)的格式](https://github.com/go-sql-driver/mysql#dsn-data-source-name)

## gorm 的基本使用 
- db, err := gorm.Open 初始化数据库连接
- db.Create 创建数据
- db.First db.Find 查询数据
- db.Model().Update 更新数据
- db.Delete 删除数据
## gorm 拓展使用
- gorm 提供了 Begin, Commit, Rollback, Transaction 用于使用事务
- gorm Hook: BeforeCreate, AfterCreate...
- gorm 默认事务功能默认开启
- 性能：可以用 SkipDefaultTransaction 关闭默认事物；使用Prepare 缓存预编译语句可以提高调用的速度
- 拓展：代码生成工具，分片库方案，手动索引，乐观锁，读写分离，OpenTelemetry扩展实现监控

## kitex 使用
- handler.go 用户在该文件里实现 IDL service 定义的方法
- 服务注册与发现已经对接了主流的服务发现与注册中心，如ETCD
- XDS 多泳道，流量路由，维护多套测试环境很方便 
- opentelemetry 可观测性
- ETCD; Nacos; Zookeeper; polaris 扩展

<br>注释:
- **etcd**是一个强一致性的分布式键值存储， 提供一种可靠的方法来存储需要由 分布式系统或计算机集群。
- **opentracing**是一种全新的开放分布式跟踪标准，适用于应用程序和 OSS 包。
- **接口描述语言**（Interface description language，缩写**IDL**），是用来描述软件组件界介面的一种计算机语言，过一种独立于编程语言的方式来描述接口。

## hertz 使用
- hertz 路由：GET; POST; PUT...
- 优先级 静态路由>命名路由>通配路由
- 参数绑定：Bind; Validate; BindAndValidate
- 代码生成工具Hz 


# 4.总结
gorm 读取操作数据库，kitex 处理远程过程调用，Hertz 对外提供api服务，它用来做接口聚合。三件套合起来提供完整的后端服务。
# 5.Ref
- [快速开始 | hertz](https://www.cloudwego.io/zh/docs/hertz/getting-started/)
- [快速开始 | kitex](https://www.cloudwego.io/zh/docs/kitex/getting-started/)
- [GORM Guides | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.cn/docs/#Install)
- Hertz 新手任务地址: <https://github.com/cloudwego/hertz/issues>
- Go 框架三件套详解(Web/RPC/ORM)实战环节-笔记服务项目地址: <br>
  优化版: <https://github.com/cloudwego/biz-demo/tree/main/easy_note>
- 普通版: <https://github.com/cloudwego/kitex-examples/tree/main/bizdemo/easy_note>
- [字节跳动开源 Go HTTP 框架 Hertz 设计实践](https://www.cloudwego.io/zh/blog/2022/06/21/%E5%AD%97%E8%8A%82%E8%B7%B3%E5%8A%A8%E5%BC%80%E6%BA%90-go-http-%E6%A1%86%E6%9E%B6-hertz-%E8%AE%BE%E8%AE%A1%E5%AE%9E%E8%B7%B5/)
 
