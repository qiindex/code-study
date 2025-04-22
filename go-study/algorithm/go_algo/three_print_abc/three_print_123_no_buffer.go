package three_print_abc

import (
	"fmt"
	"time"
)

func Main1() {
	printNumbers()
}

func printNumbers() {
	// 创建三个通道用于同步
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	// 计数器，控制循环次数
	count := 6

	// 第一个goroutine打印1
	go func() {
		for i := 0; i < count; i++ {
			<-ch1 // 等待信号
			fmt.Print("1 ")
			ch2 <- struct{}{} // 发送信号给下一个goroutine
		}
	}()

	// 第二个goroutine打印2
	go func() {
		for i := 0; i < count; i++ {
			<-ch2 // 等待信号
			fmt.Print("2 ")
			ch3 <- struct{}{} // 发送信号给下一个goroutine
		}
	}()

	// 第三个goroutine打印3
	go func() {
		for i := 0; i < count; i++ {
			<-ch3 // 等待信号
			fmt.Print("3 ")
			ch1 <- struct{}{} // 发送信号给第一个goroutine
		}
	}()

	// 启动第一轮
	ch1 <- struct{}{}

	// 等待足够时间让所有打印完成
	time.Sleep(time.Second)
}
