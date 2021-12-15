# golang版数据结与算法实现

> 数据结构与算法的资料,部分golang实现的源码

- [数据结构可视化工具usfca](https://www.cs.usfca.edu/~galles/visualization/Algorithms.html)


- 测试单个文件，一定要带上被测试的原文件

```
go test -v xxx_test.go xxx.go
```

- 测试单个方法

```
go test -v -run {函数名}
#示例
go test -v -run TestBubbleSort

#go test -v --bench=.
```

- 递归，尾递归，循环三种方式计算裴波那契数列性能对比

```
xiewj@xewj-H310MHC2:~/code/learn-dsa/other$ go test --bench=. -benchmem
goos: linux
goarch: amd64
pkg: dsa/other
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
BenchmarkFibonacci-6              446676              2694 ns/op               0 B/op          0 allocs/op
BenchmarkFibonacci2-6           48876081                24.53 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacci3-6           244700935                4.785 ns/op           0 B/op          0 allocs/op
PASS
ok      dsa/other       4.126s

```

