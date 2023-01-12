package graphStructures

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph[String]()
	g.AddVertex("A")
	start := g.Vertices["A"]
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	g.AddVertex("F")
	g.AddVertex("G")
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("A", "D")
	g.AddEdge("B", "D")
	g.AddEdge("C", "F")
	g.AddEdge("D", "E")
	g.AddEdge("E", "D")
	g.AddEdge("F", "E")
	g.AddEdge("E", "G")
	g.AddEdge("F", "G")

	visited := make(map[String]bool)
	visitation = []string{}
	g.DepthFirstSearch(start, visited)
	fmt.Println("Depth First Search:", g.String())
	visited = make(map[String]bool)
	visitation = []string{}
	g.BreadthFirstSearch(start, visited)
	fmt.Println("Breadth First Search:", g.String())
}
