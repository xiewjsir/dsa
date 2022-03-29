通过以下命令可以生成一个叫libnumber.a的静态库：
```
cd ./number
gcc -c -o number.o number.c
ar rcs libnumber.a number.o
```
在macOS和Linux系统下的gcc环境，我们可以用以下命令创建number库的的动态库：
```
cd number
gcc -shared -o libnumber.so number.c
```
CGO不仅可以使用C静态库，也可以将Go实现的函数导出为C静态库
```
go build -buildmode=c-archive -o number.a
```
CGO导出动态库的过程和静态库类似，只是将构建模式改为c-shared，输出文件名改为number.so而已：
```
go build -buildmode=c-shared -o number.so
```


```
// #cgo CFLAGS: -DPNG_DEBUG=1 -I./include
// #cgo LDFLAGS: -L/usr/local/lib -lpng
// #include <png.h>
import "C"
```
上面的代码中，CFLAGS部分，-D部分定义了宏PNG_DEBUG，值为1；-I定义了头文件包含的检索目录。LDFLAGS部分，-L指定了链接时库文件检索目录，-l指定了链接时需要链接png库。