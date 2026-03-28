package populor

import "fmt"

/*
阶乘是一个正整数的乘积，表示为 n!。
例如：5! = 5 × 4 × 3 × 2 × 1 = 120
*/
func factorial_recursion(n int) int {
	if n < 0 {
		fmt.Println("输入必须是非负整数")
		return -1
	}
	if n == 0 {
		return 1
	}
	return n * factorial_recursion(n-1)
}

/*
斐波那契数列是一个数列，其中每个数都是前两个数的和，通常以 0 和 1 开始。
例如：0, 1, 1, 2, 3, 5, 8, 13, ...
*/
func fibonacci_recursion(n int) int {
	if n < 0 {
		fmt.Println("输入必须是非负整数")
		return -1
	}
	if n < 2 {
		return n
	}
	return fibonacci_recursion(n-1) + fibonacci_recursion(n-2)
}