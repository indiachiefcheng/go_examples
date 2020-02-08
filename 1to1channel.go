/**
*@Author: chenglinguang
*@Date: 2020/2/01
*@Description: CREATE GO FILE close_channel
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	notify := make(chan int)

	datach := make(chan int, 10)

	go func() {
		<-notify
		fmt.Println("2 秒后接受到信号开始发送")
		for i := 0; i < 100; i++ {
			datach <- i
		}
		fmt.Println("发送端关闭数据通道")
		close(datach)

	}()

	time.Sleep(2 * time.Second)
	fmt.Println("开始通知发送信息")
	notify <- 1
	time.Sleep(1 * time.Second)
	fmt.Println("通知1秒后接收到数据通道数据 ")
	for {
		if i, ok := <-datach; ok {
			fmt.Println(i)

		} else {
			fmt.Println("接收不到数据中止循环")
			break
		}

	}

	time.Sleep(5 * time.Second)
}

