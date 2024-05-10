package main

import (
    "learning_go/graph"
    "fmt"
)

var discovered []int
var predecessor []int
var finished []int

func constructSampleGraph() []*graph.GraphNode {
    node_values := []int{0,1,2,3,4,5,6,7,8,9}
    graph.NewGraph()
    nodelist := make([]*graph.GraphNode, 0)
    for _, value := range node_values {
        nodelist = append(nodelist, graph.InitGraphNode(value, []*graph.GraphNode{}, "white"))
    }
    
    graph.ConnectNodes(nodelist[0], nodelist[1])
    graph.ConnectNodes(nodelist[0], nodelist[4])
    graph.ConnectNodes(nodelist[1], nodelist[2])
    graph.ConnectNodes(nodelist[2], nodelist[3])
    graph.ConnectNodes(nodelist[2], nodelist[6])
    graph.ConnectNodes(nodelist[3], nodelist[4])
    graph.ConnectNodes(nodelist[4], nodelist[5])
    
    graph.ConnectNodes(nodelist[7], nodelist[8])

    for _, node := range nodelist {
        graph.InsertGraphNode(node)
    }
    graph.PrintGraph()
    return nodelist
}

func dfs_visit(startingNode *graph.GraphNode) {
    startingNode.SetColor("gray")
    discovered[startingNode.GetValue()] += 1
    for _, neighbour := range startingNode.GetConnectedNodes() {
        if neighbour.GetColor() == "white" {
            predecessor[neighbour.GetValue()] = startingNode.GetValue()
            dfs_visit(neighbour)
        }
    }
    startingNode.SetColor("black")
    finished[startingNode.GetValue()] += 1
}
    

func main() {
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
    fmt.Printf("%v\n%v\n%v\n", discovered, finished, predecessor)
}

    
    
        
        
    
