package fibonacci

// Fibonacci 递归
func Fibonacci(n int) int {
	//if n <= 2 {
	//	return 1
	//}

	if n < 2 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Fibonacci2 尾递归 尾递归就是把当前的运算结果（或路径）放在参数里传给下层函数，不用尾递归，函数的堆栈耗用难以估量，需要保存很多中间函数的堆栈
func Fibonacci2(n, a, b int) int {
	if n == 1 {
		return a
	}

	return Fibonacci2(n-1, b, a+b)
}

// Fibonacci3 循环
func Fibonacci3(n int) int {
	var a = 1
	var b = 1

	if n <= 2 {
		return 1
	}

	for i := 2; i < n; i++ {
		a, b = b, a+b
	}

	return b
}
