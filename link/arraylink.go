package link

import "fmt"

type Value struct {
	Data string
	NextIndex int64
}

func arrayLink() {
	var link = [4]Value{}
	link[0] = Value{Data: "I",NextIndex: 2}
	link[1] = Value{Data: "you",NextIndex: 3}
	link[2] = Value{Data: "love",NextIndex: 1}
	link[3] = Value{Data: "!",NextIndex: -1}
	v := link[0]
	for {
		fmt.Printf(v.Data)
		if v.NextIndex != -1{
			v = link[v.NextIndex]
		}else {
			break
		}
	}
}
