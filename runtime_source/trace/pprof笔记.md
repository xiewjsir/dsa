#### pprof的5类profile信息

- block：goroutine的阻塞信息
- goroutine：所有goroutine的信息，下面的full goroutine stack dump是输出所有goroutine的调用栈，是goroutine的debug=2，后面会详细介绍。
- heap：堆内存的信息
- mutex：锁的信息
- threadcreate：线程信息

#### 命令行方式
```
# 下载cpu profile，默认从当前开始收集30s的cpu使用情况，需要等待30s
go tool pprof http://localhost:6060/debug/pprof/profile   # 30-second CPU profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=120     # wait 120s

# 下载heap profile
go tool pprof http://localhost:6060/debug/pprof/heap      # heap profile

# 下载goroutine profile
go tool pprof http://localhost:6060/debug/pprof/goroutine # goroutine profile

# 下载block profile
go tool pprof http://localhost:6060/debug/pprof/block     # goroutine blocking profile

# 下载mutex profile
go tool pprof http://localhost:6060/debug/pprof/mutex
```

#### 常用的：top、list、traces，分别介绍一下。
##### top 
按指标大小列出前10个函数，比如内存是按内存占用多少，CPU是按执行时间多少。
```
(pprof) top
Showing nodes accounting for 814.62MB, 100% of 814.62MB total
      flat  flat%   sum%        cum   cum%
  814.62MB   100%   100%   814.62MB   100%  main.main
         0     0%   100%   814.62MB   100%  runtime.main
```
top会列出5个统计数据：
- flat: 本函数占用的内存量。
- flat%: 本函数内存占使用中内存总量的百分比。
- sum%: 前面每一行flat百分比的和，比如第2行虽然的100% 是 100% + 0%。
- cum: 是累计量，加入main函数调用了函数f，函数f占用的内存量，也会记进来。
- cum%: 是累计量占总量的百分比。

###### list
查看某个函数的代码，以及该函数每行代码的指标信息，如果函数名不明确，会进行模糊匹配，比如list main会列出main.main和runtime.main。

```
(pprof) list main.main
Total: 1.94GB
ROUTINE ======================== main.main in /home/xiewj/code/go-other/learn-pprof/heap_demo.go
    1.94GB     1.94GB (flat, cum) 99.85% of Total
         .          .     20:   }()
         .          .     21:   
         .          .     22:   tick := time.Tick(time.Second / 100)
         .          .     23:   var buf []byte
         .          .     24:   for range tick {
    1.94GB     1.94GB     25:           buf = append(buf, make([]byte, 1024*1024)...)
         .          .     26:   }
         .          .     27:}

```
可以看到在main.main中的第25行占用了1.94GB内存，左右2个数据分别是flat和cum，含义和top中解释的一样。

##### traces
打印所有调用栈，以及调用栈的指标信息
```
(pprof) traces
File: heap_demo
Type: inuse_space
Time: Nov 29, 2021 at 10:20am (CST)
-----------+-------------------------------------------------------
     bytes:  1.94GB
    1.94GB   main.main
             runtime.main
-----------+-------------------------------------------------------
     bytes:  1.55GB
         0   main.main
             runtime.main
....... // 省略几十行
```
每个- - - - - 隔开的是一个调用栈，能看到runtime.main调用了main.main，并且main.main中占用了813.46MB内存。



> 来源：https://segmentfault.com/a/1190000019222661