package main

import "learning_go/graph"

func main() {
    graph.NewGraph()
    newHead := graph.InitGraphNode(1, []*graph.GraphNode{}, "black")
    node1, node2 := graph.InitGraphNode(2, []*graph.GraphNode{}, "black"), graph.InitGraphNode(3, []*graph.GraphNode{}, "black")
    graph.ConnectNodes(newHead, node1)
    graph.ConnectNodes(newHead, node2)
    graph.ConnectNodes(node1, node2)

    graph.InsertGraphNode(newHead)
    graph.InsertGraphNode(node1)
    graph.InsertGraphNode(node2)
    graph.PrintGraph()
}


