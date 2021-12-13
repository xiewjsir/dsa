package linknode

import "fmt"

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func CreateLink() {
	var linkNode = new(LinkNode)
	linkNode.Data = 1
	
	node1 := new(LinkNode)
	node1.Data = 2
	
	linkNode.NextNode = node1
	
	node2 := new(LinkNode)
	node2.Data = 3
	node1.NextNode = node2
	
	for {
		fmt.Println(linkNode.Data)
		if linkNode.NextNode != nil {
			linkNode = linkNode.NextNode
		} else {
			break
		}
	}
}
