package main

import "fmt"

// Go supports anonymous(匿名函数)functions, which can form closures(闭包).
// 闭包，说白了就是函数的嵌套，内层的函数可以使用外层函数的所有变量，即使外层函数已经执行完毕。

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
