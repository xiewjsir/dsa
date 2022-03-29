package main

import (
	"fmt"
	"os"
)

func main() {
	fd, err := os.Open("./syscall.go")  //一定会执行系统调用
	if err != nil {
		fmt.Println(err)
	}
	
	fd.Close()
}