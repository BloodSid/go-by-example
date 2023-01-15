这是我参与「第五届青训营 」伴学笔记创作活动的第1天

# 1. 课前准备
首先先安装 go 语言，再根据文档配置开发环境. 我使用的 IDE 是 Goland. 然后下载课程示例代码.因为示例代码中有一些 linux 相关的语句，windows 不支持，所以推荐使用 linux 作为系统环境，如果没有安装 linux 的电脑，可以使用虚拟机。我使用的是 ubuntu 20.04 系统。

```shell
git clone git@github.com:wangkechun/go-by-example.git
```
执行 go run 命令，见到 hello, world 输出说明环境配置正确

![2023-01-15 18-52-18 的屏幕截图.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/62a9ab318fad4df2a1b12633c29edd7a~tplv-k3u1fbpfcp-watermark.image?)
# 2.重点内容
- go语言的介绍
- 基础语法
- 实战

# 3.详细介绍
## go语言介绍
go 语言是谷歌出品的通用型计算机编程语言。它有高性能、高并发的编程语言，性能和 c/cpp 接近，同时 go 原生支持高并发。它进行了诸多语法简化，简单明了的语法让开发者的上手更容易。go 语言拥有极其丰富，功能完善，质量可靠的标准库，再加上完善的工具链，大大提高了开发效率和可靠性。此外，它还具有静态链接，快速编译，跨平台和垃圾回收等技术特性。
## 基础语法
- 编译和运行

```shell
# 直接运行
> go run example/01-hello/main.go
hello world
# 编译成二进制文件并运行
> go build example/01-hello/main.go
> ./main
hello world
```
- 变量的声明和初始化

```go
func main() {
    // go 可以根据上下文自动推导类型，如果有需要，也可以指定类型
    var a = "string"
    var b, c int = 1, 2
    // 初始化的简写
    d, e := true, 1.0
    // 输出以上变量
    fmt.Println(a, b, c, d, e)
}
```
- 流程控制

条件表达式不用括号，但是花括号是必须的
```go
if 7 % 2 == 0 {
     fmt.Println("even")
 } else {
     fmt.Println("odd")
 }
```
for 循环语句
```go
for i := 0; i <= 100; i++ {
   fmt.Println(i)
}
```
for 循环语句
```go
i := 0
for i < 10 {
   fmt.Println(i)
   i++
}
```
死循环，使用 break 关键字中断
```go

i := 0
for {
   fmt.Println(i)
   if i >= 10 {
      break
   }
   i++
}
```
switch-case 不需要 break, 如果需要继续执行下方的分支则使用 fallthrough 关键字
```go
switch a {
case "Daniel":
   fmt.Println("Wassup Daniel")
case "Medhi":
   fmt.Println("Wassup Medhi")
   fallthrough
case "Jenny":
   fmt.Println("Wassup Jenny")
default:
   fmt.Println("Have you no friends?")
}

```
switch 后可以不加变量，这时在 case 后加条件语句，可以让 if-else 更清晰
```go
switch {
case a < 10:
   fmt.Println("lower than ten")
case a < 100:
   fmt.Println("lower than one hundred")
default:
   fmt.Println("bigger than or equals to one hundred")
}
```
- 数组与切片
- map
- range
- 函数

go 语言中函数支持多返回值
- 指针与地址

go 语言中支持的指针操作是取地址和取值操作，不支持对指针进行加减运算
- 结构体

go 语言可以为结构体定义方法，类似于其他语言中的成员方法。定义结构体方法时，使用结构体指针就可以修改结构体
```go
type user struct {
    name string
}

func (u *user) rename(newName string) {
    u.name = newName
}
```
- 错误处理

使用函数的返回值处理异常
- 标准库：字符串-strings包；格式化字符串-Printf()方法；JSON-encoding/json包；时间-time包；数字字符串转换-strconv包；进程信息-os包
# 4.实践练习

# 5.总结

# 6.Ref
- fallthrough 的用法例子： git@github.com:GoesToEleven/GolangTraining.git