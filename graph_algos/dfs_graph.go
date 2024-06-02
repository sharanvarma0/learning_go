package main

import (
    "learning_go/graph"
    "fmt"
)

var predecessor []int

func constructSampleGraph() *graph.Graph {
    graph_map := map[int][]map[int]int{
        0: {
            map[int]int{1: 0},
            map[int]int{6: 0},
            map[int]int{8: 0},
            },
        1: {
            map[int]int{2: 0},
            map[int]int{3: 0},
        },
        2: {
            map[int]int{10: 0},
            map[int]int{11: 0},
        },
        3: {
            map[int]int{4: 0},
            map[int]int{12: 0},
        },
        4: {
            map[int]int{5: 0},
            map[int]int{13: 0},
        },
        5: {
            map[int]int{6: 0},
            map[int]int{9: 0},
        },
        6: {
            map[int]int{7: 0},
        },
        7: {
            map[int]int{8: 0},
            map[int]int{9: 0},
        },
        8: {
            map[int]int{14: 0},
        },
        9: {
            map[int]int{15: 0},
        },
    }   
    
    sampleGraph := graph.ConstructGraphFromMap(graph_map)
    sampleGraph.Print()
    return sampleGraph
}

func dfs_visit(dfsGraph *graph.Graph, startingNode *graph.Vertex) {
    startingNode.SetColor("gray")

    for _, edge := range dfsGraph.Edges[startingNode] {
        destination := edge.GetDestination()
        if destination.GetColor() != "black" && destination.GetColor() != "gray" {
            predecessor[destination.GetValue()] = startingNode.GetValue()
            dfs_visit(dfsGraph, destination)
        }
        startingNode.SetColor("black")
    }
}

func main() {
    predecessor = make([]int, 0)
    sampleGraph := constructSampleGraph()
    for range sampleGraph.Vertices {
        predecessor = append(predecessor, -1)
    }
    
    startingNode := sampleGraph.GetVertexWithValue(0)
    dfs_visit(sampleGraph, startingNode)

    // The following is only for going through disconnected nodes. In actual DFS algorithm, this step is usually
    // ignored. Thus, DFS is usually known for not being able to map disconnected graphs completely.
    // A small fix is to simply go through the graph nodes in order at the end of the first DFS call and see
    // which node is still colored white. If so, that node is disconnected from the other nodes and was not
    // Traversed in the first run.
    for _, node := range sampleGraph.Vertices {
        if node.GetColor() == "white" {
            dfs_visit(sampleGraph, node)
        }
    }
    fmt.Printf("Predecessor: %v\n", predecessor)
}

    
    
        
        
    
