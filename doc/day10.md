这是我参与「第五届青训营」伴学笔记创作活动的第10天
# 1.课前准备
预习 RPC 的基本概念和网络通讯的相关知识。
# 2.重点内容
RPC 相关的基本概念; RPC 框架的分层设计; 衡量 RPC 框架的一些核心指标; Kitex 实践
# 3.详细介绍
## 基本概念
- 相比本地函数调用：RPC 要解决的问题：1.函数映射；2.数据流转换成字节流；3.网络传输
- RPC 的理论模型由5部分组成：User; User-Stub; RPC-Runtime; Server-Stub; Server
- IDL 通过一种中立的方式描述接口，使得不同平台的程序可以互相通信。在IDL文件的基础上，可以通过工具生成代码。
- Caller & 生成代码 -> 编码 -> 协议层 -> 网络通信 -> 协议层 -> 解码 -> Callee & 生成代码
- 使用RPC 的**好处**：1.单一职责，有利于分工和运维；2.可扩展性强；3.故障隔离；服务整体的可扩展性更高。**问题**：对方宕机怎么处理；网络异常怎么处理；请求积压怎么处理。
## RPC 框架分层设计
1. 编解码层：客户端和服务端依赖同一份IDL文件，生成不同语言的GenCode。**数据格式**: java.io.Serializable; JSON, XML, CSV 等文本格式；BinaryProtocol, Protobfuf 等二进制编码。选型：兼容性；通用性；性能。
2. 协议层：消息切分：特殊结束符；变长协议：length+body。协议构造例子： Thrift 的 [THeader](https://github.com/apache/thrift/blob/master/doc/specs/HeaderFormat.md) 协议
3. 网络通信：网路库封装Socket API 提供易用API,并提供连接管理，事件分发，优雅退出和异常处理等能力。网络库中基于 IO 多路复用系统调用实现的 Poll 的意义在于将可读/可写状态通知和实际文件操作分开，并支持多个文件描述符通过一个系统调用监听以提升性能。网络库的核心功能就是去同时监听大量的文件描述符的状态变化(通过操作系统调用)，并对于不同状态变更，高效，安全地进行对应的文件操作。

## RPC 框架核心指标
- 稳定性：保障策略：熔断，限流，超时控制。请求成功率：负载均衡，重试。长尾请求：BackupRequest

- 易用性：开箱即用 + 周边工具

- 扩展性：Middleware，Option；核心层是支持扩展的，代码生成工具也支持插件

- 观测性：三件套：Log、Metric 和 Tracing；内置观测性服务，用于观察框架内部状态，  当前环境变量、配置参数、缓存信息、内置 pprof 服务用于排查问题

- 高性能：连接池和多路复用：复用连接，减少频繁建联带来的开销；高性能编解码协议：Thrift、Protobuf、Flatbuffer 和 Cap'n Proto 等；高性能网络库：Netpoll 和 Netty 等
## kitex实践
1.  框架文档 [Kitex](https://www.cloudwego.io/zh/docs/kitex/)
2.  自研网络库 [Netpoll](https://www.cloudwego.io/zh/docs/netpoll/)，背景：原生库无法感知连接状态；原生库存在 goroutine 暴涨的风险
3.  扩展性：支持多协议，也支持灵活的自定义协议扩展
4.  性能优化，参考 [字节跳动 Go RPC 框架 KiteX 性能优化实践](https://www.infoq.cn/article/spasfyqgaaid5rguinl4)
- 网络优化：1.调度优化；2.LinkBuffer 减少内存拷贝，从而减少 GC；3.引入内存池和对象池
- 编解码优化：1.Codegen：预计算提前分配内存，inline，SIMD等；2.JIT：无生产代码，将编译过程移到了程序的加载（或首次解析）阶段，可以一次性编译生成对应的 codec 并高效执行
5.  合并部署: 微服务过微，引入的额外的传输和序列化开销越来越大；将强依赖的服务统计部署，有效减少资源消耗

# 4.总结
RPC(远程函数调用）在后端开发中是相当重要的基础技术。RPC框架分为解编码层；协议层；网络传输层。RPC 的核心指标包括稳定性，可扩展性和高性能等。理解和熟练掌握RPC技术的使用在开发中可以起到很大的帮助。

# 5.Ref
- [RPC 框架分层设计 - 掘金 (juejin.cn)](https://juejin.cn/course/bytetech/7142811324462923783/section/7142809631831228429)
- [‌深入浅出 RPC 框架 副本.pptx - 飞书云文档 (feishu.cn)](https://bytedance.feishu.cn/file/boxcn5DUtKdJDDitx8NHShv2xZd)
- [ RPC 原理与实践 - 掘金 (juejin.cn)](https://juejin.cn/post/7196322025114779703#heading-33)
- 官方文档 [Kitex](https://www.cloudwego.io/zh/docs/kitex/)； [Netpoll](https://www.cloudwego.io/zh/docs/netpoll/)
- [字节跳动 Go RPC 框架 KiteX 性能优化实践_架构_字节跳动技术团队_InfoQ精选文章](https://www.infoq.cn/article/spasfyqgaaid5rguinl4)
-  [字节跳动微服务架构体系演进_架构_字节跳动技术团队_InfoQ精选文章](https://www.infoq.cn/article/asgjevrm8islszo7ixzh)