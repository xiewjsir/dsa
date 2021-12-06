- 测试单个文件，一定要带上被测试的原文件

```
go test -v xxx_test.go xxx.go
```

- 测试单个方法

```
go test -v -test.run {函数名}
#示例
go test -v -test.run TestBubbleSort
```
