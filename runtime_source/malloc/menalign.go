package main


/*
对于任意类型的变量 x ，unsafe.Alignof(x) 至少为 1。
对于 struct 结构体类型的变量 x，计算 x 每一个字段 f 的 unsafe.Alignof(x.f)，unsafe.Alignof(x) 等于其中的最大值。
对于 array 数组类型的变量 x，unsafe.Alignof(x) 等于构成数组的元素类型的对齐倍数。
 */

import (
	"bytes"
	"fmt"
	"strings"
	"unsafe"
)

type Args struct {
	num1 int
	num2 int
}

type Flag struct {
	num1 int16
	num2 int32
}

type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

func main() {
	var a [5]int
	fmt.Println(unsafe.Alignof(a),unsafe.Sizeof(a))
	fmt.Println(unsafe.Alignof(Args{}),unsafe.Sizeof(Args{}))
	fmt.Println(unsafe.Alignof(Flag{}),unsafe.Sizeof(Flag{}))
	
	fmt.Println(unsafe.Sizeof(demo1{})) // 8
	fmt.Println(unsafe.Sizeof(demo2{})) // 12
	
	sb := strings.Builder{}
	sb.WriteString("aaaa")
	fmt.Println(sb.String())
	
	buf := new(bytes.Buffer)
	buf.WriteString("ccc")
	fmt.Println(buf.String())
}