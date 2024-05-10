package graph

import "fmt"

var GlobalGraph []*GraphNode

func NewGraph() {
    GlobalGraph = make([]*GraphNode, 0)
}

func PrintGraph() {
    for _, node := range GlobalGraph {
        fmt.Printf("%d -> ", node.GetValue())
        for _, connectedNode := range node.GetConnectedNodes() {
            fmt.Printf("%d, ", connectedNode.GetValue())
        }
        fmt.Printf("\n")
    }
}

func InsertGraphNode(newNode *GraphNode) bool {
    GlobalGraph = append(GlobalGraph, newNode)
    return true
}

func ConnectNodes(node1, node2 *GraphNode) bool {
    node1.ConnectNode(node2)
    node2.ConnectNode(node1)
    return true
}

func InitGraphNode(value int, connectedNodes []*GraphNode, color string) *GraphNode {
    return &GraphNode{
        value: value,
        connectedNodes: connectedNodes,
        color: color,
    }
}

func GetNodeWithValue(value int) *GraphNode {
    for _, node := range GlobalGraph {
        if node.value == value {
            return node
        }
    }
    return nil
}
