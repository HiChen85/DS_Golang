package graph

import (
	"fmt"
)

type Set map[string]struct{}

func NewSet() *Set {
	set := new(Set)
	*set = make(map[string]struct{})
	return set
}

func (s *Set) String() string {
	var str = "{"
	for k, _ := range *s {
		str += fmt.Sprintf("%v, ", k)
	}
	str += "}"
	return str
}

func (s *Set) Add(n string) bool {
	if s.Has(n) {
		return false
	} else {
		(*s)[n] = struct{}{}
		return true
	}
}

func (s *Set) Remove(n string) {
	delete(*s, n)
}

func (s *Set) Has(n string) bool {
	if _, ok := (*s)[n]; ok {
		return true
	} else {
		return false
	}
}

func (s *Set) IsEmpty() bool {
	return 0 == len(*s)
}

func (s *Set) Len() int {
	return len(*s)
}

func (s *Set) toSlice() []string {
	var sli []string
	for k, _ := range *s {
		sli = append(sli, k)
	}
	return sli
}

type Graph struct {
	Nodes *Set
	Edges map[string][]string // key-value[i]表示 key 到 value[i] 有边,而 value[i] 到key 不一定有边
}

func NewGraph() *Graph {
	g := new(Graph)
	g.Nodes = NewSet()
	g.Edges = make(map[string][]string)
	return g
}

func (g *Graph) String() string {
	nodes := "Nodes:{"
	for k, _ := range *g.Nodes {
		k += ", "
		nodes += k
	}
	nodes += "} \n"

	edges := "Edges:{"
	for k, v := range g.Edges {
		s := "(" + k + fmt.Sprintf("->%v", v) + "), "
		edges += s
	}
	edges += "} \n"

	return nodes + edges

}

func (g *Graph) IsEmpty() bool {
	if g.Nodes.Len() == 0 && len(g.Edges) == 0 {
		return true
	} else {
		return false
	}
}

func (g *Graph) AddNode(val string) bool {
	if g.Nodes.Add(val) {
		g.Edges[val] = []string{}
		fmt.Println("已成功添加")
		return true
	} else {
		return false
	}
}

// AddEdge : the edge is from node1 to node2
func (g *Graph) AddEdge(node1, node2 string) bool {
	if g.Nodes.Has(node1) {
		for _, v := range g.Edges[node1] {
			if v == node2 {
				fmt.Println("该条边已存在")
				return false
			}
		}
		g.AddNode(node2)
		g.Edges[node1] = append(g.Edges[node1], node2)
		return true
	} else {
		g.AddNode(node1)
		g.AddNode(node2)
		g.Edges[node1] = append(g.Edges[node1], node2)
		return true
	}
}

func (g *Graph) AddNodeFrom(nodes []string) {
	for _, v := range nodes {
		if !g.AddNode(v) {
			fmt.Println(v, "已存在")
		}
	}
}

type Path struct {
	Visited bool
	Dist    int
	adj     string
}

func BFS(g *Graph, src string) {
	if g.IsEmpty() {
		fmt.Println("空图")
		return
	}
	if !g.Nodes.Has(src) {
		return
	}
	// 此处利用 带缓冲区的 channel 来实现队列
	q := make(chan string, g.Nodes.Len())
	pathList := make(map[string]Path)

	q <- src
	for len(q) != 0 {
		n := <-q
		if !pathList[n].Visited {
			fmt.Printf(n, " ")
			pathList[n] = Path{Visited: true}
			for _, v := range g.Edges[n] {
				q <- v
			}
		}
	}
	fmt.Println()
}

func DFS(g *Graph, src string) {
	if g.IsEmpty() {
		fmt.Println("空图")
		return
	}
	if !g.Nodes.Has(src) {
		return
	}

	s := make([]string, 0)
	pathList := make(map[string]bool)
	// 入栈
	s = append(s, src)
	for len(s) != 0 {
		top := len(s) - 1
		e := s[top]
		if !pathList[e] {
			fmt.Println(e)
			pathList[e] = true
		}

		if len(g.Edges[e]) != 0 {
			i := 0
			for _, v := range g.Edges[e] {
				if !pathList[v] {
					s = append(s, v)
					break
				}
				i++
			}
			// 当所有边都被访问过,则出栈
			if i == len(g.Edges[e]) {
				s = s[:top]
			}
		} else { // 当图节点没有边时停止.
			s = s[:top]
		}

	}

}
