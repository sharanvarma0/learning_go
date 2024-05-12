package main

import (
    "learning_go/graph"
    "fmt"
)

var distance []int
var predecessor []int
var queue []*graph.GraphNode

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

func bfs_visit() {
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        for _, neighbour := range node.GetConnectedNodes() {
            if neighbour.GetColor() != "white" {
                continue
            }
            predecessor[neighbour.GetValue()] = node.GetValue()
            distance[neighbour.GetValue()] = distance[node.GetValue()] + 1
            neighbour.SetColor("gray")
            queue = append(queue, neighbour)
        }
        
        node.SetColor("black")
    }
}

func main() {
    distance, predecessor = make([]int, 0), make([]int, 0)
    queue = make([]*graph.GraphNode, 0)
    node_list := constructSampleGraph()
    for range node_list {
        distance = append(distance, -1)
        predecessor = append(predecessor, -1)
    }
    startingNode := node_list[0]
    distance[startingNode.GetValue()] = 0
    startingNode.SetColor("gray")
    queue = append(queue, startingNode)
    bfs_visit()

    fmt.Printf("Queue: %v\n Distance: %v\n Prececessors: %v\n", queue, distance, predecessor)
}

    
    
        
        
    
