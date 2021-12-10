# learn-dsa
 > 数据结构与算法的资料及部分实现的源码

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
