package dbfs

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	fmt.Println("GoLang, Graph DFS and BFS implementation")
	fmt.Println("DFS : Depth First Search")
	fmt.Println("BFS : Breadth First Search")
	
	g := NewGraph()
	
	g.AddVertex("ajinkya")
	g.AddVertex("francesc")
	g.AddVertex("manish")
	g.AddVertex("albert")
	
	g.AddEdge("albert", "ajinkya")
	g.AddEdge("ajinkya", "albert")
	g.AddEdge("francesc", "ajinkya")
	g.AddEdge("francesc", "manish")
	g.AddEdge("manish", "francesc")
	g.AddEdge("manish", "albert")
	
	g.DFS("francesc")
	g.CreatePath("francesc", "albert")
}