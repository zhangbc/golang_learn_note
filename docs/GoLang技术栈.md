# GoLang基础

> 学习资料：https://golang-tech-stack.com/tutorial/golang/intro
>
> https://www.runoob.com/go/go-tutorial.html
>
> https://www.bilibili.com/video/BV1zR4y1t7Wj?p=1&vd_source=616bf6a004def119a1b27e570687e7ea

## 1，基本知识

#### 1）goLang基本概念

`Go`（又称 `Golang`）是 `Google` 的 `Robert Griesemer`，`Rob Pike` 及 `Ken Thompson` 开发的一种静态强类型、编译型语言。`Go` 语言语法与 `C` 相近，但功能上有：内存安全，`GC`（垃圾回收），结构形态及 `CSP-style` 并发计算。

- `go` 语言特点

> 1）背靠大厂，`google` 背书，可靠
>
> 2）天生支持并发（最显著特点）
>
> 3）语法简单，容易上手
>
> 4）内置 `runtime`，支持垃圾回收
>
> 5）可直接编译成机器码，不依赖其他库
>
> 6）丰富的标准库
>
> 7）跨平台编译

`Go` 语言被设计成一门应用于搭载 `Web` 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。

对于高性能分布式系统领域而言，`Go` 语言无疑比大多数其它语言有着更高的开发效率。

#### 2）`Go` 基础安装及配置

> 官网地址：https://go.dev/

```bash
# 查看版本
☁  ~  /usr/local/go/bin/go version
go version go1.19.1 darwin/amd64

# 配置环境变量
☁  ~  vi .zshrc
86 export PATH=$PATH:/usr/local/go/bin
:wq!

# 配置go环境
☁  ~  go env -w GO111MODULE="on"
☁  ~  go env -w GOPROXY="http://goproxy.cn"
```

#### 3）`Go` 基础知识

- `Go` 常用命令 

> `build`: 编译包和依赖
>
> `clean`: 移除对象文件
>
> `doc`: 显示包或者符号的文档
>
> `env`: 打印 `go` 的环境信息
>
> `bug`: 启动错误报告
>
> `fix`: 运行 `go tool fix`
>
> `fmt`: 运行 `gofmt` 进行格式化
>
> `generate`: 从 `processing source` 生成 `go` 文件
>
> `get`: 下载并安装包和依赖
>
> `install`：编译并安装包和依赖
>
> `list`：列出包
>
> `run`：编译并运行 `go` 程序
>
> `test`：运行测试
>
> `tool`：运行 `go` 提供的工具
>
> `version`：显示 `go` 的版本
>
> `vet`: 运行 `go tool vet`

```bash
# 初始化项目
☁  go_learning_note  go mod init go_learning_note
# 下载包
☁  go_learning_note  go get github.com/go-sql-driver/mysql
# 编译
☁  golang_learn_note [go_basic_note] ⚡  go build main.go
# 运行
☁  golang_learn_note [go_basic_note] ⚡  go run main.go  
# 清除对象
☁  golang_learn_note [go_basic_note] ⚡  go clean 
```

- 标识符

标识符的英文是`identifier`，通俗的讲，就是给变量、常量、函数、方法、结构体、数组、切片、接口起名字。

标识符的组成

> 1）标识符由数字、字母和下划线(`_`)组成。`123 abc _`
>
> 2）只能以字母和下划线(`_`) 开头。`abc123 _sysVar 123abc`
>
> 3）标识符区分大小写。 `name Name NAME`

- `Go` 语言关键字	

> `break、default、func、interface、select、case、defer、go、map、struct、chan、else、goto`
>
> `package、switch、const、fallthrough、if、range、type、continue、for、import、return、var`
>

- `Go` 语言命名规范

> 1）任何需要对外暴露的名字必须以大写字母开头；不需要对外暴露的则应该以小写字母开头；
>
> 2）当命名（包括常量、变量、类型、函数名、结构字段等）以一个大写字母开头时，那么使用这种形式的标识符的对象就可以被外部包的代码所使用，这被称为 **导出**；命名如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且是可用的；
>
> 3）包名称：保持 `package` 的名字和目录保持一致，尽量采取有意义的包名称，简短，有意义，尽量和标准库不要冲突；包名应该为小写单词，不要使用下划线或混合大小写；
>
> 4）文件命名：尽量采取有意义的文件名，简短，有意义，应该为小写单词，使用下划线分割各个单词；
>
> 5）结构体命名：采用驼峰命名法，首字母根据访问控制大写或小写；`struct` 申明和初始化格式采用多行；
>
> 6）接口命名：命名规则基本和结构体命名一样，单个函数的结构以“`er`” 作为后缀，如 `Reader`、`Writer`；
>
> 7）变量命名：和结构体命名类似，一般遵循驼峰法，首字母根据访问控制大写或小写，但是遇到特有名词时，需要遵循以下规则：
>
> > 如果变量为私有，且特有名词为首个单词，则使用小写。
>
> 8）常量命名：常量均需使用全部大写字母组成，并使用下划线分词，如果是枚举类型的常量，则需要先创建相应类型；
>
> 9）错误处理：其原则是不能丢弃任何有返回的 `err` 的调用，不要使用 `_`丢弃，必须全部处理；
>
> 10）单元测试：单元测试文件命名规范为 `example_test.go`，测试用例的函数名称必须以 `Test` 开头。	


- 数据类型

| 序号 | 类型和描述                                                   |
| ---- | ------------------------------------------------------------ |
| 1    | **布尔型** 布尔型的值只可以是常量 `true` 或者 `false`。一个简单的例子：`var b bool = true`。 |
| 2    | **数字类型** 整型 `int` 和浮点型 `float32`、`float64`，`Go` 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。 |
| 3    | **字符串类型:** 字符串就是一串固定长度的字符连接起来的字符序列。`Go` 的字符串是由单个字节连接起来的。`Go` 语言的字符串的字节使用 `UTF-8` 编码标识 `Unicode` 文本。 |
| 4    | **派生类型:** 包括：`(a)` 指针类型（`Pointer`）`(b)` 数组类型 `(c)` 结构化类型(`struct`) `(d) Channel` 类型 `(e)` 函数类型 `(f)` 切片类型 `(g)` 接口类型（`interface`） `(h) Map` 类型 |

- `golang` 变量

变量是计算机语言中能**储存**计算结果或能表示值的抽象概念。不同的变量保存的**数据类型**可能会不一样。

声明变量

`Go` 语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。 并且 `Go` 语言的变量声明后**必须使用**。

声明变量的语法

```go
var identifier type
```

> `var`：声明变量关键字
>
> `identifier`：变量名称
>
> `type`：变量类型

变量声明

> 1）指定变量类型，如果没有初始化，则变量默认为零值；
>
> 2）根据值自行判定变量类型；
>
> 3）如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误（短变量声明）
>
> > 该方法仅限于函数内部，函数外部不能使用。

```go
// 批量声明变量
var (
    name string
    pAge int
    isMan bool
)
```

变量初始化

```go
var 变量名 类型 = 表达式
```

类型推导

```go
var pName = "zhangbc"
```

初始化多个变量

```go
var name, site, age = "wangwu", "www.baidu.com", 30
```

短变量声明

> 在函数内部，可以使用 `:=` 运算符对变量进行声明和初始化；函数外部不能使用。

```go
name := "zhangbc"
```

匿名变量

> 如果我们接收到多个变量，有一些变量使用不到，可以使用下划线 `_`  表示变量名称，这种变量叫做 **匿名变量**。

```go
// 匿名变量
name, _ = getNameAndAge()
fmt.Printf("name: %v\n", name)
```


- `golang` 常量

常量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：

```go
const identifier [type] = value
```

`iota`

`iota`，特殊常量，可以认为是一个可以被编译器修改的常量。

`iota` 在 `const` 关键字出现时将被重置为 `0` (`const` 内部的第一行之前)，`const` 中每新增一行常量声明将使 `iota` 计数一次(`iota` 可理解为 `const` 语句块中的行索引)。

```go
// iota
const (
    a1 = iota  // 0
    a2 = iota  // 1
    a3 = iota  // 2
)
```

使用 `_` 跳过某些值

```go
// iota
const (
    a1 = iota  // 0
    _
    a3 = iota  // 2
)
```

`iota` 声明中间插队

```
  // iota
  const (
      a1 = iota  // 0
      a2 = 100
      a3 = iota  // 2
  )
```

- `Go` 语言布尔类型

	 `go` 语言中布尔类型有两个常量值：`true` 和 `false` ，布尔类型经常用于条件判断语句，或循环语句，或用于逻辑表达式中。		
>
> **注意**：不能使用 0 和 非0 表示真假 

```go
age := 8
condition := age >= 18
if condition {
    fmt.Println("你已成年了.")
} else {
    fmt.Println("你还为成年.")
}
```

- `Go`语言数字类型

`Go` 语言支持整型和浮点型数字，并且原生支持复数，其中位的运算采用补码。

`Go` 也有基于架构的类型，例如：`int`、`uint` 和 `uintptr`。

> `int`、`uint` 在32位操作系统上，均使用32位（4个字节）；在64位操作系统上，均使用64位（8个字节）；
>
> `uintptr` 的长度被设定为足够存放一个指针即可。

| 序号 | 类型和描述                                                   |
| ---- | ------------------------------------------------------------ |
| 1    | **uint8** 无符号 8 位整型 (0 到 255)                         |
| 2    | **uint16** 无符号 16 位整型 (0 到 65535)                     |
| 3    | **uint32** 无符号 32 位整型 (0 到 4294967295)                |
| 4    | **uint64** 无符号 64 位整型 (0 到 18446744073709551615)      |
| 5    | **int8** 有符号 8 位整型 (-128 到 127)                       |
| 6    | **int16** 有符号 16 位整型 (-32768 到 32767)                 |
| 7    | **int32** 有符号 32 位整型 (-2147483648 到 2147483647)       |
| 8    | **int64** 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

浮点型

| 序号 | 类型和描述                        |
| ---- | --------------------------------- |
| 1    | **float32** IEEE-754 32位浮点型数 |
| 2    | **float64** IEEE-754 64位浮点型数 |
| 3    | **complex64** 32 位实数和虚数     |
| 4    | **complex128** 64 位实数和虚数    |

其他数字类型

以下列出了其他更多的数字类型：

| 序号 | 类型和描述                                 |
| ---- | ------------------------------------------ |
| 1    | **`byte`** 类似 `uint8`                    |
| 2    | **`rune`** 类似 `int32`                    |
| 3    | **`uint`** 32 或 64 位                     |
| 4    | **`int`** 与 `uint` 一样大小               |
| 5    | **`uintptr`** 无符号整型，用于存放一个指针 |

- `GoLang` 字符串

一个 `Go` 语言字符串是一个任意**字节的常量序列**。

在 `Go` 语言中，字符串字面量使用双引号 `""` 或反引号 `'` 来创建。

> 双引号 `""` ：创建可解析的字符串，支持转义，但不能用来引用多行；
>
> 反引号 `'` ：创建原生的字符串字面量，可能由多行组成，但不支持转义。	

转义字符

> `Go` 语言的字符串常见转义字符包括回车，换行，单双引号，制表符，反斜杠等。

字符串的切片操作

```go
// 字符串的切片操作
str_split := "Hello World"
start := 2
end := 5
fmt.Printf("str_split[start]: %c\n", str_split[start])
fmt.Printf("str_split[start:end]: %v\n", str_split[start:end])
fmt.Printf("str_split[start:]: %v\n", str_split[start:])
```

字符串函数

```go
// 字符串函数
fmt.Printf("len(str_split): %v\n", len(str_split))
fmt.Printf("strings.Split(str_split, \" \"): %v\n", strings.Split(str_split, " "))
fmt.Printf("strings.Contains(str_join, \"Hello\"): %v\n", strings.Contains(str_join, "Hello"))
fmt.Printf("strings.ToLower(str_split): %v\n", strings.ToLower(str_split))
fmt.Printf("strings.ToUpper(str_split): %v\n", strings.ToUpper(str_split))
fmt.Printf("strings.HasPrefix(str_split, \"Hello\"): %v\n", strings.HasPrefix(str_split, "Hello"))
fmt.Printf("strings.HasSuffix(str_split, \"world\"): %v\n", strings.HasSuffix(str_split, "world"))
fmt.Printf("strings.Index(str_split, \"ll\"): %v\n", strings.Index(str_split, "ll"))
```

字符串格式化输出

```go
// 格式化输出
site := WebSite{Name: "douke360"}
fmt.Printf("site: %v\n", site)
fmt.Printf("site: %#v\n", site)
fmt.Printf("site: %T\n", site)
```

- `GoLang` 运算符

1. 算术运算符

> +（相加），-（相减），*（相乘），/（相除），%（求余）
>
> ++（自增），--（自减）在 `Go` 语言中说单独的语句，并不是运算符

2. 关系运算符

> ==、!=、>、>=、<、<=

3. 逻辑运算符

> &&、||、!

4. 位运算符

> &、|、^、>>（右移，除）、<<（左移，乘）

5. 赋值运算符

> =、+=、-=、*=、/=、%=、<<=、>>=、&=、|=、^=

- `Go` 语言中的流程控制

  **条件语句** 是用来判断给定的条件是否满足（表达式是否为 `true` 或 `false`），并根据判断的结果（真或假）决定执行的语句。

  > **`if` 语句** ：由一个布尔表达式后紧跟一个或多个语句组成；
  >
  > > 初始变量可以声明在布尔表达式里面，注意它的作用域；
  > >
  > > 不需使用括号将条件包含起来；
  > >
  > > 大括号必须存在，即使只有一行语句；
  > >
  > > 左括号必须在 `if` 或 `else` 的同一行；
  > >
  > > 在 `if` 之后，条件语句之前，可以添加变量初始化语句，使用 `:` 进行分隔
  >
  > **`if...else` 语句** ：`if` 语句后使用可选的 `else` 语句，`else` 语句中的表达式在布尔表达式为 `false` 时执行；
  >
  > **`if` 嵌套语句**：可以在 `if` 或 `else if` 语句中嵌入一个或多个 `if` 或 `else if` 语句；
  >
  > **`switch` 语句**：用于基于不同条件表达式执行不同动作；
  >
  > > 支持多条件匹配；
  > >
  > > 不同的 `case` 之间不使用 `break` 分隔，默认只会执行一个 `case`；
  > >
  > > 如果想要执行多个 `case`，需要使用 `fallthrough` 关键字，也可以使用 `break` 终止；
  > >
  > > 分支可以使用**表达式**，如 `a > 18`
  >
  > **`select` 语句**：类似于 `switch` 语句，但是 `select` 语句会随机执行一个可运行的 `case`，如果没有 `case` 可运行，它将阻塞，直到有 `case` 可运行。  

`Go` 语言中的循环只有 `for` 循环，去除了 `while`、`do while` 循环，使用起来更加简洁。

```go
// 永真循环
for {
		fmt.Printf("for: %v\n", "run....")
	}
```

`Go` 语言中可以使用 `for range` 遍历数组、切片、字符串、`map` 及通道（`channel`），通过 `for range` 遍历的返回值有如下规律：

> 数组、切片、字符串返回索引和值；
>
> `map` 返回键和值；
>
> 通道（`channel`）只返回通道内的值

流程控制关键字：`break、continue、goto`

> `break`：可以结束 `for`、`switch`、`select` 的代码块
>
> > 单独在 `select` 语句中使用 `break` 和不使用 `break` 没有区别；
> >
> > 单独在表达式 `switch` 语句并且没有 `fallthough` 中使用 `break` 和不使用 `break` 没有区别；
> >
> > 单独在表达式 `switch` 语句并且有 `fallthough` 中使用 `break` 能够终止`fallthough` 后面的 `case` 语句的执行；
> >
> > 带标签的 `break`，可以跳出多层 `select /switch` 作用域
>
> `continue`：只能用在循环中，在 `go` 中只能用在 `for` 循环中，可以终止本次循环，进行下一次循环；在`continue` 语句后添加标签时，表示开始标签对应的循环。
>
> `goto`：通过标签进行代码间的 **无条件跳转**，在快速跳出循环、避免重复退出上有一定的帮助。

- `Go` 数组

> **数组** 是相同据类型的一组数据的集合，数组一旦定义长度不能变，数组可以通过下标（或叫索引）来访问元素。

```go
// 数组定义
var variable_name [SIZE] variable_type 
```

- `Go` 切片

> **切片（Slice）** 是一个拥有相同类型元素的可变长度的序列；可以理解为：可变长度的数组，其底层就是使用数组实现的，增加了自动扩容功能。

```go
//切片定义
var identifier []type
var slice []type = make([]type, len)
slice := make([]type, len)
make([]T, length, capacity)
```

**空（nil）切片**：一个切片在未初始化之前默认为 `nil`, 长度为0，容量为0。

- `Go map`

> `map` 是一种 `key:value`键值对的数据结构容器，其内部实现时哈希表（`hash`），是一种引用类型。

```go
// 定义
var map_variable map[key_data_type] value_data_type
map_variable = make(map[key_data_type] value_data_type)
```





## 2，工具使用

1）`gopm`

> https://github.com/gpmgo/docs/blob/master/zh-CN/Quickstart.md

```bash
# 安装：默认的会存放到GOBIN，如果没有配置%GOBIN%环境变量，那么会默认安装到%GOPATH%下的bin目录
☁  crawler [main] ⚡  go get -u github.com/gpmgo/gopm
##############解决无法download github问题###############
# https://ipaddress.com/website/github.com
# 140.82.113.3
# https://ipaddress.com/website/fastly.net#ipinfo(github.global.ssl.fastly.net)
# 151.101.1.6 151.101.65.6 151.101.129.6 151.101.193.6
###################################################
140.82.113.3    github.com
151.101.1.6     github.global.ssl.fastly.net
151.101.65.6    github.global.ssl.fastly.net
151.101.129.6   github.global.ssl.fastly.net
151.101.193.6   github.global.ssl.fastly.net
###################################################
# 刷新DNS
☁  ~  sudo killall -HUP mDNSResponder;say DNS cache has been flushed
# 设置环境变量
☁  golang  vi ~/.zshrc
 87 export PATH=$PATH:/Users/zhangbc/go/bin
☁  golang  source ~/.zshrc

# 用gopm get -g代替go get
# 不采用-g参数，会把依赖包下载.vendor目录下面；
# 采用-g 参数，可以把依赖包下载到GOPATH目录中；
```



# GoLang 项目实战

## 1，Go 分布式爬虫

> 课程地址：https://www.bilibili.com/video/BV1bq4y1Z7p2?p=1

#### 1，分布式系统介绍

- 分布式系统

> 多个节点
>
> > 容错性
> >
> > 可扩展性（性能）
> >
> > 固有分布性
>
> 消息传递
>
> > 节点具有私有存储
> >
> > 易于开发
> >
> > 可扩展性（功能）
> >
> > 对比：并行计算
> >
> > 传递方法：
> >
> > > `REST` ：用户（网页、app、API客户端）与服务器之间，对外
> > >
> > > `RPC`：服务器与服务器之间，模块内部
> > >
> > > 中间件：模块之间
>
> 完成特定需求

- 分布式架构 `vs` 微服务架构

> 分布式：指导节点之间如何通信
>
> 微服务：鼓励按照业务划分模块
>
> 微服务架构通过分布式架构来实现

- 多层架构 `vs` 微服务架构

> 微服务架构具有更多的“服务”
>
> 微服务通常需要配合自动化测试，部署，服务发现等
>
> 目前我们倾向于微服务架构

#### 2，分布式爬虫架构

![arch_go](images/go_pro_20221007004416.png)

目前的问题：

> 限流问题
>
> > 单节点能够承受的流量有限
> >
> > 将 `Worker` 放到不同的节点
>
> ![image-20221007010245554](images/image-20221007010245554.png)
>
> 去重问题
>
> > 单节点能够承受的去重数据量有限
> >
> > 无法保存之前的去重结果
> >
> > 基于 `Key-Value Store` （如 `Redis`）进行分布式去重
>
> ![image-20221007010111210](images/image-20221007010111210.png)
>
> ![image-20221007010514238](images/image-20221007010514238.png)
>
> 数据存储问题
>
> > 存储部分的结构，技术栈和爬虫部分区别很大
> >
> > 进一步优化需要特殊的 `ElasticSearch` 技术背景
> >
> > 固有分布式
>
> ![image-20221007010847210](images/image-20221007010847210.png)

![image-20221007011051150](images/image-20221007011051150.png)

# 参考资料

- [国内的go get问题的解决 --gopm](https://www.jianshu.com/p/db9e6ae0d227)
- [GitHub无法访问、443 Operation timed out的解决办法](https://juejin.cn/post/6844904193170341896)