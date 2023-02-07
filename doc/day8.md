这是我参与「第五届青训营」伴学笔记创作活动的第8天

# 1.课前准备

预习相关概念

# 2.重点内容

常见的分布式系统；系统模型；理论基础；分布式事务；共识协议；分布式实践 为什么TCP采用三次握手而不是两次和四次？

# 3.详细介绍
- 分布式：分布式系统是计算机程序的集合，这些程序利用跨多个独立计算节点的计算资源来实现共同的目标。 **优势**：去中心化，低成本，弹性，资源共享，可靠性高；**挑战**：普遍的节点故障；不可靠的网络；异构的硬件；安全
- Why : 数据爆炸，对存储和计算有大规模运用的诉求；成本低；帮助学习者理解后台服务器之间协作的机理
- How: 掌握分布式理论；了解一致性协议
- What : 理清负载，规模，一致性要求和稳定性要求，制定技术方案

## 常见的分布式系统
- 分布式存储： Google File System; Ceph; Hadoop HDFS; Zookeeper
- 分布式数据库：Google Spanner; TiDB; HBase; MongoDB
- 分布式计算： Hadoop; Spark; YARN

## 系统模型
- 故障模型：Byzantine failure 拜占庭故障；ADB，前者的特例，节点可以篡改数据，但是不能伪造其他节点的数据；Performance failure; Omission failure; Crash failure; Fail-stop failure。**故障四维度**：正确性；时间；状态；原因
- 拜占庭将军问题：两支军队理论上永远无法达成共识，不存在理论解。共识与消息传递的不同：即使保证了消息传递成功，也不能保证达成共识。TCP 三次握手是在两个方向确认包的序列号，增加了超时重试，是两将军问题的一个工程解。
- 多将军问题：当有3m+1个将军，其中 m 个“叛徒”时。可以增加 m 轮协商最终达成一致。比特币是拜占庭容错的系统，其他分布式系统都不是。
- 共识和一致性：Eventually consistent 最终一致性。linearizability 线性一致性，（强一致性）：一旦某个客户端读取到新值，所有客户端都必须返回新值。保证线性一致性，多个节点之间势必要进行协商，系统可用性会受损。
- 时间和事件顺序：定义 happened before 关系，用 "->" 表示。当且仅当 a 不-> b 且 b 不-> a 时，两个事件a和b是并发的(concurrent)。 Lamport 逻辑时钟: 利用 tick line 将时间分成原子；利用逻辑时钟，可以对整个系统中的事件进行全序排序。

## 理论基础
- CAP理论：百分之百的一致性，可用性，分区容错性三者无法兼得。在网络发生分区的情况下，必须在可用性和一致性之间做出选择；近似解决办法：把故障节点的负载转移给备用节点负责，可用性和一致性都有所损失，但不至于完全失去其中一半。
- ACID理论：事务是数据库管理系统执行过程中的一个逻辑单元，它能保证一个事务中的所有操作要么全部执行，要么全部不执行。它有4个特性 ACID: 原子性，一致性，隔离性，持久性。
- BASE 理论是对 CAP 中一致性和可用性权衡的结果，核心思想: 基本可用，软状态，最终一致性。

## 分布式事务
- 两阶段提交：prepare 阶段，commit 阶段 1. 引入协调者和参与者，互相进行网络通信；2. 所有节点都采用预写式日志，且日志写入后被保持在可靠的存储设备上；3.所有节点不会永久性损坏，即使损坏也可恢复。协调者和参与者都宕机的情况下，无法确认状态，需要介入，防止数据库进入不一致的状态。**问题**：性能问题；协调者单调故障问题；**网络分区带来的数据不一致**。
- 三阶段提交：将两阶段提交中的 prepare 阶段拆成 canCommit 和 preCommit。解决了：1.单点故障问题；2.阻塞问题。另外引入超时机制，在等待超时之后，会继续进行事务的提交。没解决：1.性能问题；2.网络分区场景带来的数据一致性问题。
- MVCC: 一种并发控制的方法，维持一个数据的多个版本使读写操作没有冲突。所以既不会阻塞写，也不阻塞读。MVCC为每个修改保存一个版本，和事务的时间戳相关联。可以提高并发性能，解决脏读的问题。时间戳的实现：硬件方案，物理时钟；时间戳预言机(TSO)。


## 共识协议
- Quoruw NWR 模型：为了保证强一致性需要保证 W+R>N；Quoruw NWR 模型将CAP的选择交给用户，是一种简化的一致性模型。数据允许被覆盖的时可能有并发更新问题。适合不允许数据被覆盖的场景。
- Raft 协议：一种分布式一致性算法。角色：Leader, Follower, Candidate。精髓：Term 任期号，单调递增，每个 Term 内最多只有一个 Leader。为了避免双主的情况出现需要：election timeout > lease timeout, 新 leader 上任，自从上次心跳之后一定超过了 election timeout, 旧 leader 大概率能够发现自己的 lease 过期。
- Paxos 协议：Multi-Paxos 可以并发修改日志，而 Raft 写日志必须是连续的。Multi-Paxos 可以随机选主，不必最新最全的节点当Leader。优势：写入并发性能高，所有节点都能写入；劣势：没有一个节点有完整的最新的数据，恢复流程复杂，需要同步历史记录。


# 4.总结
分布式系统的核心是一致性，可用性和分区容错性之间的权衡取舍。对于不同的取舍有不同的实现方案，对应不同的应用场景。分布式理论中的共识协议复杂而巧妙，通过对这些原理的深入理解，可以帮助我们理解分布式系统中互相协作的机理，从而正确地运用分布式系统。

# 5.Ref

1.  1978年Leslie Lamport发表《*Time, Clocks, and the Ordering of Events in a Distributed System》*
2.  [RFC 793: Transmission Control Protocol (rfc-editor.org)](https://link.juejin.cn?target=https%3A%2F%2Fwww.rfc-editor.org%2Frfc%2Frfc793 "https://www.rfc-editor.org/rfc/rfc793")
3. [Transmission Control Protocol - Wikipedia](https://link.juejin.cn?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FTransmission_Control_Protocol "https://en.wikipedia.org/wiki/Transmission_Control_Protocol")
4.   [Byzantine fault - Wikipedia](https://link.juejin.cn?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FByzantine_fault "https://en.wikipedia.org/wiki/Byzantine_fault")

  


