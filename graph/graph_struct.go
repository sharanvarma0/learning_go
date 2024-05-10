package graph

type GraphNode struct {
    value int
    color string
    connectedNodes []*GraphNode
}

type Graph interface {
    GetConnectedNodes() []*GraphNode
    GetValue() int
}


func (g *GraphNode) GetConnectedNodes() (nodes []*GraphNode) {
    nodes = g.connectedNodes
    return
}

func (g *GraphNode) GetValue() (val int) {
    val = g.value
    return
}

func (g *GraphNode) ConnectNode(n *GraphNode) bool {
    g.connectedNodes = append(g.connectedNodes, n)
    return true
}

func (g *GraphNode) GetColor() string {
    return g.color
}

func (g *GraphNode) SetColor(color string) {
    g.color = color
}
