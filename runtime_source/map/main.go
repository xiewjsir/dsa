package main

import (
	"strconv"
)

func main() {
	m := make(map[string]string, 16*6.5+1)
	for i := 0; i < 500; i++ {
		m["key"+strconv.Itoa(i)] = "val" + strconv.Itoa(i)
	}

	//fmt.Println(m["key0"])
}
