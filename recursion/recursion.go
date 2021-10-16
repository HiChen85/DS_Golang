package recursion

import "fmt"

func CountDown(n int) {
	if n <= 0 {
		return
	}
	fmt.Printf("%v...\n", n)
	CountDown(n - 1)
	fmt.Println("当前的函数栈", n)
}
