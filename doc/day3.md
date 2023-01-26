这是我参与「第五届青训营」伴学笔记创作活动的第3天

# 1.课前准备
- 下载示例工程代码https://github.com/Moonlight-Zhao/go-project-example
- 预习并发，单元测试等知识
# 2.重点内容
- go工程实践:concurrence包;并发编程
- test包:单元测试;mock
- benchmark包:基准测试
# 3.详细介绍
## 并发
在并发场景下，Go可以充分发挥多核优势高效运行，这是Go取得成功的一个关键因素。go的并发使用的是协程，协程与线程的主要区别是，协程在用户态，线程在内核态。协程比线程更轻量，切换代价更低。
```go
    go func(){
        // 函数体
    }()
```
go 中使用 go 关键字发起goroutine实现高并发, go关键字后跟 函数名（参数列表） 发起协程。也可以像上文调用匿名函数。<br>
go 提倡通过通信共享内存，而不是通过共享内存实现通信。前者例如协程之间通过管道进行通信实现内存中数据的共享，而后者则例如维护一个临界区通过锁机制实现内存数据的共享从而实现协程之间的通信。<br>
### channel
channel 是用来传递数据的数据结构
```go
// 通道的声明
ch := make(chan 元素类型, [缓冲大小])
```
带缓冲区的通道允许发送端的数据发送和接受端的数据处于异步状态。一个通道可以通过 range 关键字方便地读取输入其中的数据，在通道关闭之前，如果通道没有数据，那么接收端协程就会阻塞以等待新数据或调用 close(ch)
```go
func main(){
    ch := make(chan int)
    go func() {
        // 函数执行完后会调用 defer 关闭通道
        defer close(ch)
        for i := 0; i < 10; i++{
            // <- 表示把结果放入通道
            ch <- i * i
        }
    for j := range ch {
        // 循环读取 ch 直到通道关闭
        println(j)
    }
}
```
这里漏写close 会导致所有协程都进入休眠 -死锁，程序会非正常终止。

### lock & waitGroup
go 也支持用锁进行协程之间的同步
```go
// 锁的声明
var lock sync.Mutex
// 上锁和解锁
lock.Lock()
lock.Unlock()

// wg 的声明
var wg sync.WaitGroup
// 增加计数器
wg.add(5)
// 计数器减一
wg.Done()
// 直到计数器归零前，阻塞当前协程
wg.Wait()
```
## 依赖管理
go 的依赖管理分为三代
1. **GOPATH**：用GOPATH/src 目录保存依赖项目的源码，go get 下载最新的包到该目录下。**缺点**：不能实现包的多版本控制
2. **Vendor**: 在项目目录下增加 vendor 目录，把项目的依赖包的副本存在该目录中。在这种机制下，vendor中的依赖优先使用，如果vendor没有才去GOPATH中找。这样就解决了多个项目需要一个包的不同版本的冲突问题。**缺点**：无法控制依赖的版本。在一个项目中依赖同一个包的不同版本时会出现冲突。
3. **Go Module**: 通过 go.mod 文件管理依赖包版本，通过指令工具管理依赖包。**缺点**：通过代码托管平台获取依赖包源码，因为代码托管平台可能进行删改，则构建稳定性和依赖可用性无法保证，也会增加第三方平台的压力。**解决方案**：使用 go proxy 管理依赖分发。
### Go Module 依赖配置
- 依赖标识：\[module Path]\[version/pseudo-version]
- 版本标识 MAJOR.MINOR.PATCH， MAJOR 不同不兼容，MINOR 不同向后兼容。
- 伪版本：vX.0.0-时间-哈希；包括基础版本前缀，时间戳，和commit对应的12位哈希前缀
- indirect标识符：表示间接依赖
- incompatible后缀：一个依赖的新的主版本要建立一个新目录（如v3)并用不同的go.mod文件管理，来表明不同主版本之间的不兼容性。而为了兼容历史仓库，对于没有单独go.mod文件且主版本在2或以上的依赖需要加上此后缀以让go module 按照不同的模块来处理。
- 版本规则：在有多个兼容版本可选时，选择最低的兼容版本
- go proxy: 从服务站点获取缓存的源站点内容，从而保证了“不可修改的”和“可用的”依赖分发。GOPROXY 保存代理站点的url列表，并用“direct”表示源站，站点url之间逗号分割
### 工具
- go get: 语法go get example.org/pkg@version 获取指定包的制定版本。**注意**go get 在go1.17以前的版本会在下载源代码后编译和安装可执行程序。新版本则只做下载。
- go mod: 用法：go mod init 初始化，创建go.mod文件；go mod download 下载模块到本地缓存。go mod tidy 增加需要的依赖，删除不需要的依赖。

## 测试
### 单元测试
- 测试文件使用 \_test.go 结尾，测试函数名 func TestXxx(\*testing.T)
- 初始化逻辑放到 TestMain 中
- 断言：利用 testify 包提供的 assert.Equals 方法进行结果的断言
- 运行： go test \[flags] \[packages]
### 覆盖率测试
go test \[文件] --cover
### Mock
打桩测试：利用 monkey 包提供的 Patch 和 Unpatch 方法进行打桩。原理：在运行时通过 unsafe 包，将内存中函数的地址替换为代替函数的地址，这样调用原函数时会跳转到代替函数。**作用**：隔离被测试单元的外部依赖，保证测试的稳定性和幂等性。
### 基准测试
用于对代码进行性能分析，指令： go test -bench=函数名。测试函数: func BenchmarkXxx（b \*testing.B)。
例如go原生的 rand 在多线程并行的情况下会劣化，经过测试发现可以用字节开发的 fastrand 代替。

## 分层结构模型
- repository 数据层：外部数据的增删改查，如数据库的数据，封装外部的数据操作，对逻辑层透明
- service 逻辑层：处理核心业务逻辑,计算打包业务实体entity并供给视图层
- controller 视图层：处理和外部的交互逻辑，以view 对象形式返回给客户端。
# 4.总结
- go 为并发提供了协程这一强大的工具，正确使用go roution 可以充分发挥多核处理器的优势。
- go 最新的依赖管理方法是 go modules，其对应的依赖管理三要素分别是 1. 配置文件，依赖描述： go.mod；2. 中心仓库管理依赖库 go proxy; 3. 本地工具 go get/mod
- 单元测试是开发阶段，开发者对代码的“单元”进行的功能验证。衡量测试套件的完备性的一个指标是代码覆盖率。设计单元测试时，要求测试分支相互独立，全面覆盖，同时也促进开发者把测试单元的粒度设计的更小，促进了单一职责设计原则的实行。这样就有助于提高代码覆盖率。
- 项目实践中应用了需求分析，用例图，ER-图等技术，设计了数据层，逻辑层，视图层的分层结构。利用 Gin 框架提供http服务实现需求。
# 5.Ref
-   锁Lock [pkg.go.dev/sync](https://pkg.go.dev/sync)
-   线程同步WaitGroup [pkg.go.dev/sync](https://pkg.go.dev/sync)
-   Go Module : [go.dev/blog/using-…](https://go.dev/blog/using-go-modules")
-   单元测试概念和规则：[go.dev/doc/tutoria…](https://go.dev/doc/tutorial/add-a-test%EF%BC%9Bhttps://pkg.go.dev/testing)
-   Mock测试：[github.com/bouk/monkey](https://github.com/bouk/monkey)
-   基准测试：[pkg.go.dev/testing#hdr…](https://pkg.go.dev/testing#hdr-Benchmarks)

-   web框架：Gin - [github.com/gin-gonic/g…](https://github.com/gin-gonic/gin#quick-start)

-   分层结构设计：[github.com/bxcodec/go-…](https://github.com/bxcodec/go-clean-arch)

-   文件操作：读文件[pkg.go.dev/io](https://pkg.go.dev/io)
-   数据查询：索引[www.baike.com/wikiid/5527…](https://www.baike.com/wikiid/5527083834876297305?prd=result_list&view_id=5di0ak8h3ag000)

  


 
