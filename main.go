// 1(无缓冲区的channel)
//package main
//
//import "fmt"
//
//func main() {
//	c := make(chan string)
//	go func() {
//		fmt.Println("waiting for information...")
//		s := <-c
//		fmt.Println("received information:", s)
//	}()
//
//	c <- "hello"
//	fmt.Println("main_goroutine finished")
//}

// 2(带有缓冲区的channel的读写)
//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	ch := make(chan int, 5)
//	wait := sync.WaitGroup{}
//	wait.Add(1)
//	go func() {
//		defer wait.Done()
//		ch <- 1
//		ch <- 2
//		ch <- 3
//		//fmt.Println("阻塞恢复")
//	}()
//	wait.Wait()
//	ch <- 5
//	ch <- 6
//	num := <-ch
//	fmt.Println("num:", num)
//	close(ch)
//	for data := range ch {
//		fmt.Println("data:", data)
//	}
//}

// 3 工作池（simple）
package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Do func(id int)
}

func main() {
	channel := make(chan Job, 10)
	waiter := sync.WaitGroup{}
	waiter.Add(15)
	for i := range 4 {
		go func(id int) {
			for job := range channel {
				job.Do(id)
				waiter.Done()
			}
		}(i) //这里必须要让i作为值拷贝，传入这个匿名函数  相关知识点：闭包陷阱
	}
	for i := range 15 {
		j := i
		job := Job{
			Do: func(id int) {
				fmt.Printf("【工作编号：%d，执行任务%d】\n", id, j)
			},
		}
		channel <- job
	}
	close(channel)
	waiter.Wait()
}
