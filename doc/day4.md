这是我参与「第五届青训营」伴学笔记创作活动的第4天
# 1.课前准备
-   获取代码 [github.com/wolfogre/go…](https://github.com/wolfogre/go-pprof-practice) 到本地
-   尝试使用 test 命令，和 -bench 参数
# 2.重点内容
性能分析工具 pprof 的使用
# 3.详细介绍
## 常见编码规范
原则：简单性。可读性。生产力。代码被写出来花费的时间远远长于代码被阅读的时间，所以为了提高团队的效率，代码一定要易读。
- 使用 gofmt 和 goimports 自动格式化代码
- 注释：应该解释代码作用；解释代码如何做的；解释为什么要这样做；解释什么情况下会有异常
- 公共符号使用要注释
- 命名: 变量使用的范围越大，则越需要携带更多的上下文信息。函数名不携带包名信息，因为函数名和包名往往同时出现。
- 控制流程：尽量降低嵌套层数，保证流程清晰，方便维护。
## 错误和异常处理
- 简单错误，（只出现一次的错误，且不需要被捕获），使用 errows.New 和 fmt.Errorf 来直接表示简单错误
- Warp 和 Unwarp。对于复杂的错误，需要对错误进行包装，生成一个 error 的跟踪链。在fmt.Errorf中使用 %w 将错误关联进错误链中。使用 errors.Is 判断错误链上是否含有特定错误。使用 errors.As 从错误链上获取特定错误。
- panic: 在程序发生无法处理或继续执行也没有意义的错误时，使用panic使程序终止。建议：业务代码中不使用 panic；在程序启动时，发生不可逆转的错误时使用 panic。
- recover: 可以在触发宕机时让程序恢复过来，执行完对应的 defer 后，从宕机点退出当前函数后继续执行。常用于panic时，退出前进行的后处理时使用，例如打印 debug.Stack 中的调用栈信息后再退出。只能在当前 goroutine 的被 defer 的函数中生效。
## 性能优化建议
使用 基准测试 benchmark 工具，以基准测试的结果等数据进行优化
- make 初始化切片/ map 时，尽可能提供容量信息，减少申请空间的次数
- 大切片上 reslice 小切片，造成大切片不释放：用 copy 代替 reslice
- 字符串拼接使用 strings.Builder
- 在语义上要求时，使用空接口体节省内存 如：map\[int]struct{}
- atomic 包和 Mutex: 前者通过硬件实现，效率比锁高，后者通过操作系统实现，属于系统调用。前者用来保护一个变量，后者保护一段逻辑

## 性能优化--pprof
pprof 工具用于可视化和分析性能分析数据。**采样**：可以对 CPU，堆内存，goroutine，锁，阻塞调用和系统线程的数据进行采集。**展示**：可以通过列表，热点-调用图，火焰图，源码，反汇编等视图展示采集到的性能指标

- 采样数据生成：1. 文件：可以在 go test 命令使用 -cpuinfo cpu.out 标志，程序运行结束后生成分析文件。2. http：对于持续运行的程序，可以导入net/http/pprof包并提供http服务向外提供采样接口。
- 采样数据查看：使用 go tool pprof \[ file | url ] 在命令行进行查看。此命令使用 -http 标志可在网页查看。
- flat 当前函数本身的执行耗时；cum 当前函数加上其调用的总耗时。当 flat=cum 时，说明函数没有调用其他函数；flat=0 时说明函数中只有其他函数的调用。

### 采样原理
- **CPU**： 采样记录所有的调用栈和它们的占用时间。原理：每秒暂停一百次，记录当前的调用栈信息，通过记录到的次数推断函数的运行时间。同时启动一个写缓冲的goroutine, 每隔100ms读取已经记录的信息，写入输出流。
- **堆内存**：记录堆上内存的分配，默认每分配 512KB 采样一次。采样率可以修改。与CPU不同，内存采样是持续的。采样指标： inuse = alloc - free，正在使用 = 累计分配 - 累计释放
- **goroutine**: 记录所有用户发起且在运行中的（包括 runtime.main）goroutine的调用栈信息。
- **threadCreate**: 记录程序创建的所有系统线程的信息。它和goroutine的实现非常相似，都是立刻触发的全量记录，都在Stop the world后，遍历所有goroutine/线程的列表并输出，最后继续运行。
- **block**: 记录采样操作的次数和耗时。阻塞耗时超过阈值的才会被记录，1为每次阻塞均记录。
- **Mutex**: 记录争抢锁的次数和耗时。只记录固定比例的锁操作，1为每次加锁均记录。它和block的实现一样，都是“主动上报”，在block 或 mutex操作发生时，计算消耗时间，连同调用栈一起主动上报给采样器。
# 4.总结
注释的原则是补充和指明代码未表达出的上下文信息，不能缺少，也不能冗余重复。命名的原则是降低理解成本，需要考虑上下文信息，设计简洁恰当的名称。error 应该尽可能提供简明的上下文信息，方便定位问题。而 panic 则用于真正异常的情况。<br>
性能优化：不要过早优化，不要过度优化，在保证正确和代码质量的前提出下提高程序性能。go 提供了强大的性能分析工具 pprof，提供了丰富的信息和多样的视图帮助开发者定位性能问题。<br>
性能调优通用流程：1.建立性能评估手段；2.分析性能数据，定位性能瓶颈；3.重点优化项改造；4.优化效果验证。一些典型的性能问题原因：使用库不规范；高并发场景优化不足；服务的重复调用。<br>
总之，性能调优的原则是要依靠数据而不是猜测。性能调优要在保证正确性的基础上，定位优化主要瓶颈。
# 5.Ref
[go 官方代码](https://github.com/golang/go/tree/master/src)<br>
[Pungyeon/clean-go-article: A reference for the Go community that covers the fundamentals of writing clean code and discusses concrete refactoring examples specific to Go. (github.com)](https://github.com/Pungyeon/clean-go-article)<br>
[Effective Go - The Go Programming Language](https://go.dev/doc/effective_go)<br>
[Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)