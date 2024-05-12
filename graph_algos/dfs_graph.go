package main

import (
    "learning_go/graph"
    "fmt"
)

var discovered []int
var predecessor []int
var finished []int
var counter int

func constructSampleGraph() []*graph.GraphNode {
    node_values := []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
    graph.NewGraph()
    nodelist := make([]*graph.GraphNode, 0)
    for _, value := range node_values {
        nodelist = append(nodelist, graph.InitGraphNode(value, []*graph.GraphNode{}, "white"))
    }
    
    graph.ConnectNodes(nodelist[0], nodelist[1])
    graph.ConnectNodes(nodelist[0], nodelist[6])
    graph.ConnectNodes(nodelist[0], nodelist[8])
    graph.ConnectNodes(nodelist[1], nodelist[2])
    graph.ConnectNodes(nodelist[1], nodelist[3])
    graph.ConnectNodes(nodelist[2], nodelist[10])
    graph.ConnectNodes(nodelist[2], nodelist[11])
    graph.ConnectNodes(nodelist[3], nodelist[4])
    graph.ConnectNodes(nodelist[3], nodelist[12])
    graph.ConnectNodes(nodelist[4], nodelist[5])
    graph.ConnectNodes(nodelist[4], nodelist[13])
    graph.ConnectNodes(nodelist[5], nodelist[6])
    graph.ConnectNodes(nodelist[5], nodelist[9])
    graph.ConnectNodes(nodelist[6], nodelist[7])
    graph.ConnectNodes(nodelist[7], nodelist[8])
    graph.ConnectNodes(nodelist[7], nodelist[9])
    graph.ConnectNodes(nodelist[8], nodelist[14])
    graph.ConnectNodes(nodelist[9], nodelist[15])
    
    for _, node := range nodelist {
        graph.InsertGraphNode(node)
    }
    graph.PrintGraph()
    return nodelist
}

func dfs_visit(startingNode *graph.GraphNode) {
    startingNode.SetColor("gray")
    counter += 1
    discovered[startingNode.GetValue()] = counter
    for _, neighbour := range startingNode.GetConnectedNodes() {
        if neighbour.GetColor() == "white" {
            predecessor[neighbour.GetValue()] = startingNode.GetValue()
            dfs_visit(neighbour)
        }
    }
    startingNode.SetColor("black")
    counter += 1
    finished[startingNode.GetValue()] = counter
}
    

func main() {
    counter = 0
    discovered, finished, predecessor = make([]int, 0), make([]int, 0), make([]int, 0)
    node_list := constructSampleGraph()
    for range node_list {
        discovered = append(discovered, -1)
        finished = append(finished, -1)
        predecessor = append(predecessor, -1)
    }
    
    dfs_visit(node_list[0])
    for _, node := range node_list {
        if node.GetColor() == "white" {
            dfs_visit(node)
        }
    }
    fmt.Printf("Discovered: %v\nFinished: %v\nPredecessor: %v\n", discovered, finished, predecessor)
}

    
    
        
        
    
