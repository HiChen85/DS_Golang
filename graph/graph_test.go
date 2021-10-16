package graph

import "testing"

func TestGraph(t *testing.T) {
	g := NewGraph()

	g.AddNode("1")
	g.AddNode("2")
	t.Log(g)
	g.AddNodeFrom([]string{"1", "2", "3", "4", "5"})
	t.Log(g)
	g.AddEdge("1", "2")
	g.AddEdge("1", "4")
	g.AddEdge("2", "3")
	g.AddEdge("2", "5")
	g.AddEdge("3", "4")
	g.AddEdge("3", "5")
	g.AddEdge("3", "1")
	g.AddEdge("4", "5")
	t.Log(g)

	//BFS(g, "1")
	DFS(g, "2")
}
