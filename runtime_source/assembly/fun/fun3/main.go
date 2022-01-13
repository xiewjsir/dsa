package main

func main() {
	printsum(1, 2)
}

func printsum(a, b int) {
	var ret = sum(a, b)
	println(ret)
}

func sum(a, b int) int {
	return a+b
}
