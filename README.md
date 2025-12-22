# Lanshan-Go-2025-Homework
蓝山
//Lv0
package main

import "fmt"

func main() {
	fmt.Println("Hello,SheckHall")
}

//Lv1
package main

import "fmt"

const pi = 3.14

func main() {
	var r = 5
	area := pi * float64(r) * float64(r)
	fmt.Println("area:", area)
}

//Lv2
package main

import "fmt"

func main() {
	var sum int
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//Lv3
package main

import "fmt"

func cac(n int) int {
	if n == 0 {
		return 1
	}
	var result int = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(cac(n))
}

// LvX
package main

import "fmt"

func average(sum int, count int) float64 {
	return float64(sum) / float64(count)
}
func main() {
	var sum int = 0
	var num int = 0
	for {
		var n int
		fmt.Print("请输入一个整数（输入0结束）: ")
		fmt.Scan(&n)
		if n == 0 {
			break
		} else {
			num += 1
			sum += n
		}
	}
	fmt.Printf("平均成绩为%.2f,", average(sum, num))
	if average(sum, num) >= 60 {
		fmt.Println("成绩合格")
	} else {
		fmt.Println("成绩不合格")
	}
}
