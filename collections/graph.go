package collections

type Graph struct {
	Nodes map[string]map[string]bool
}

// 创建一个新的图
func NewGraph() *Graph {
	g := &Graph{
		Nodes: make(map[string]map[string]bool),
	}
	return g
}

// 添加顶点
func (g *Graph) AddNode(node string) {
	if _, found := g.Nodes[node]; !found {
		g.Nodes[node] = make(map[string]bool)
	}
}

// 添加边
func (g *Graph) AddEdge(from, to string) {
	g.AddNode(from)
	g.AddNode(to)
	g.Nodes[from][to] = true
	g.Nodes[to][from] = true // 因为是无向图，所以也添加反向边
}

// 检查两个顶点之间是否有边
func (g *Graph) HasEdge(from, to string) bool {
	if _, found := g.Nodes[from]; !found {
		return false
	}
	return g.Nodes[from][to]
}
