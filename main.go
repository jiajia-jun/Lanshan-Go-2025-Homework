package main

import (
	"Goland/cmd/goTeam5/homework1/pool"
	"fmt"
	"sync"
)

var result1, result2 int64 // 全局变量结果
var lock sync.Mutex        // 全局互斥锁

// 定义累加操作Task函数
func Task1() {
	lock.Lock() // 上锁，实现并发操作
	defer lock.Unlock()
	for range 100 {
		result1++
	}
}
func Task2() {
	lock.Lock() // 上锁，实现并发操作
	defer lock.Unlock()
	for range 50 {
		result2++
	}
}
func main() {
	var tasknum int
	var goroutines int
	fmt.Print("输入累加任务数目：")
	fmt.Scan(&tasknum)
	fmt.Print("输入协程数目：")
	fmt.Scan(&goroutines)

	mypool1 := pool.NewPool()
	mypool1.Put(Task1, tasknum)
	mypool1.Arrange(goroutines)
	mypool1.Wait()
	fmt.Println("Task1:", result1)

	fmt.Println()

	mypool2 := pool.NewPool()
	mypool2.Put(Task2, tasknum)
	mypool2.Arrange(goroutines)
	mypool2.Wait()
	fmt.Println("Task2:", result2)
}
