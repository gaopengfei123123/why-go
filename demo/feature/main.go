package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// 创建一个缓冲
	ch := make(chan string, 3)
	go loopPro(true, ch)
	go loopPro(false, ch)
	fmt.Println("start")

	for i := 0; i < 20; i++ {
		fmt.Print(<-ch)
	}

}

// 函数多个返回值
func multipartReturn(a int, b int) (res int, err error) {
	res = a + b
	if res > 3 {
		return res, nil
	}
	return 0, errors.New("math: square root of negative number")
}

// 延迟执行
func deferDemo() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("输出：", i)
	}
}

// 用于演示 goroute 的 demo
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
