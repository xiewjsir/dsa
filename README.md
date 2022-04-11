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

```go
b->c,初始,c为白
a->c,c赋值a,触发写屏障,把c置为灰
b->nil,b删除对c的引用时,触发写屏障,把c置为灰

// 灰色赋值器 Dijkstra 插入屏障 触发点在写入时
//slot 当前下游对像
//ptr新下游对象
func DijkstraWritePointer(slot *unsafe.Pointer, ptr unsafe.Pointer) {
	shade(ptr)
	*slot = ptr
}

// 黑色赋值器 Yuasa 删除屏障 触发点在删除时
func YuasaWritePointer(slot *unsafe.Pointer, ptr unsafe.Pointer) {
	shade(*slot)
	*slot = ptr
}

// 混合写屏障
func HybridWritePointerSimple(slot *unsafe.Pointer, ptr unsafe.Pointer) {
	shade(*slot)
	shade(ptr)
	*slot = ptr
}
```

- [Go语言高级编程(Advanced Go Programming)](https://chai2010.cn/advanced-go-programming-book/ch1-basic/readme.html)
- [Go 语言设计与实现](https://www.bookstack.cn/read/draveness-golang/30cd28c181a56e61.md#abr725)
- [go语言调度器源代码情景分析](https://mp.weixin.qq.com/mp/homepage?__biz=MzU1OTg5NDkzOA==&hid=1&sn=8fc2b63f53559bc0cee292ce629c4788&scene=18&devicetype=iOS15.0&version=1800112b&lang=zh_CN&nettype=WIFI&ascene=7&session_us=gh_8b5b60477260&fontScale=100&pass_ticket=zibsUD4Cw79t%2B7Tq1GkyJp7Zp0Q8%2BI0hol4Q9zIy4QPiaps6nyJLOGf0v51IGzNt&wx_header=3&scene=1)
- [数据结构和算法（Golang实现）](https://www.bookstack.cn/read/hunterhug-goa.c/algorithm-sort-quick_sort.md)