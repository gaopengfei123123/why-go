
## 大体了解一下语法

### 安装

官网安装包地址: https://golang.org/dl/
以 centos 为例
```bash
$ wget https://dl.google.com/go/go1.4.linux-amd64.tar.gz
$ tar -C /usr/local -zxvf go1.4.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin/
$ go
```
GOROOT 不用管它...
GOPATH 目录结构
```
bin/                                # 生成的执行文件                
pkg/                                # 编译时用到的外部库执行文件
src/                                # 我们开发的各种库和 go get 到的各种库源码
```


### Hello world

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}
```
运行 `go run main.go`



### 资源类型

常用: 
- `int/int8/int16/int32/int64`
- `uint/uint8/uint16/uint32/uint64`
- `float32/float64`
- `string`
- `array`
- `slice`(切片)
- `map`(字典)
- `point`(指针)

面向对象: 
- `struct`
- `interface`

特色: 
- `chan`


### 变量赋值
```go
var a int
a = 1
// 或
a := 1

// 数组
var arr [10]int

// 切片
var slice1 []int = make([]int, 3)
var slice1 := make([]int, 3, 10)    // 预留了10字节内存
slice1 = []int{1, 2, 3}
slice2 := slice1[2:]		// [3]

// 接口
var any interface{}         // 空接口接一切
any = slice2				// [3]

// 字典
var mp map[string]string
mp["a"] = "诶"
mp["b"] = "必"
// 同上 , 不过遍历时是无序的
mp := map[string]string{"a": "aaa", "b": "bbb"}
// 这样写得清楚这是干什么的
var mp map[string]interface{}
```

### 逻辑处理

```go
// for循环
for i := 0;i < 5 ; i++ {
    fmt.Println("输出：",i)
}


// 遍历数组/字典
mp := map[string]int{"a": 1, "b": 2}
for index, value := range mp {
	fmt.Printf("输出：index: %s, value: %d \n", index, value)
}

// 条件判断
if i > 1 {
    ....
}

// 隐式执行 break
switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
}
```


### 面向对象

#### struct

```go
package main

import (
	"fmt"
)

type human struct {
	Name string
	Age  int
}

// 隐式的继承 human 的属性
type ming struct {
	human
	Skill string
}

// 添加结构体所属的方法
func (h human) Say() {
	fmt.Printf("我叫%s,今年%d", h.Name, h.Age)
}

func main() {
	var xm ming
	xm.Name = "小明"
	xm.Age = 18
	xm.Skill = "跑得快"
	fmt.Println(xm)

	xm.Say()
}

```


#### interface

> 当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子。 -- [鸭子类型](https://zh.wikipedia.org/wiki/%E9%B8%AD%E5%AD%90%E7%B1%BB%E5%9E%8B)

```go
package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type PHPProgrammer struct {
}

func (p PHPProgrammer) Speak() string {
	return "PHP is the best language! "
}

type PassBy struct {
}

func main() {
	animals := []Animal{Dog{}, Cat{}, PHPProgrammer{}}
	// animals[3] = PassBy{}		// 编译不会通过, 因为没有实现 Animal 接口
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
```


--- 

## go 的特色

### defer 延迟处理
```go
func OpenFile() bool{
  file.Open("file/path")

  if condition1 {
    file.Close()
    return false
  }

  if condition2 {
    file.Close()
    return false
  }

  file.Close()
  return true
}
```

对比一下
```go
func OpenFile() bool{
  file.Open("file/path")
  defer file.Close()

  if condition1 {
    return false
  }

  if condition2 {
    return false
  }
  return true
}
```


```go
...
func main () {
  for i :=0;i < 5 ; i++{
    defer fmt.Println("输出：",i)
  }
}
...
// 输出： 4
// 输出： 3
// 输出： 2
// 输出： 1
// 输出： 0
```

### 多返回值
```go
func multipartReturn(a int, b int) (res int, err error) {
	res = a + b
	if res > 3 {
		return res, nil
	}
	return 0, errors.New("math: square root of negative number")
}
```

### goroutinue 并发
```go
func main() {
	go loop(true)
	loop(false)
}

func loop(desc bool) {
	if desc {
		for i := 10; i > 0; i-- {
			fmt.Printf("loop1: %d \n", i)
		}
	} else {
		for j := 0; j < 10; j++ {
			fmt.Printf("loop2: %d \n", j)
		}
	}
}
```


### channel 多线程通信
```go

func main() {
	// 创建一个缓冲
	ch := make(chan string, 3)
	go loopPro(true, ch)
	go loopPro(false, ch)
	for i := 0; i < 20; i++ {
		fmt.Print(<-ch)
	}
}

// 加入 channel 的概念
func loopPro(desc bool, ch chan<- string) {
	if desc {
		for i := 10; i > 0; i-- {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("loop1: %d \n", i)
		}
	} else {
		for j := 0; j < 10; j++ {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("loop2: %d \n", j)
		}
	}
}
```


### 跨平台
`http-server`的 demo
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("route: http://localhost:9090/")
	fmt.Println("route: http://localhost:9090/hello")

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", pageHandler)

	http.ListenAndServe(":9090", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi baby"))
}
```

可根编译成指定平台的执行文件
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

### 调用 c 代码
```go
package main

/* 下面是 C 代码 */

// int add(int a, int b) {
//     return a + b;
// }
import "C"

import "fmt"

func main() {
	a := C.int(1)
	b := C.int(2)
	value := C.add(a, b)
	fmt.Printf("%v\n", value)
}

```

### 支持webAssembly

### 支持RPC



## 总结一下

### go 的特性
* 开发效率和运行的效率都在可接受范围
* 成熟的官方库和活跃的社区, 还有官方出品的工具链
* 语法简单(25个关键字), 对于初学者来说不用受配置环境的苦
* 良好兼容 c 语言
* 语言层面支持并发
* 运行时对外部的依赖极少 (一个golang-app 约等于 nginx+php+redis)


### go 相比 php 有哪些优势?
* go 的部署简单, 运行时只要把一个二进制文件扔到机器上就行, 这一点在容器化方面优势很大
* 可控的多线程和更简单的多线程间通信
* 编译型语言对动态语言天生的性能优势, 变量在开发过程中不会突然从 str 变成 array
* 微服务相关的项目多
* 自 go 1.5 以后, 底层完全采用的 go 自己的代码重构

### go 较于 php 的劣势
* 编译造成的繁琐操作
* 代码量偏多, 个人体验大概是php的2倍左右
* 需要有多线程的意识, 和资源作用域的意识
* 字典类型要考虑并发读写安全的问题, 需要有`锁`的概念, 这一点 swoole 也有类似的情况


[代码来源](http://bastengao.com/blog/2017/12/go-cgo-c.html)

--- 
