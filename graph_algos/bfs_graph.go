package main

import (
    "learning_go/graph"
    "fmt"
)

var distance []int
var predecessor []int
var queue []*graph.Vertex

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
    
    newGraph := graph.ConstructGraphFromMap(graph_map)
    newGraph.Print()
    return newGraph
}

func bfs_visit(graphObj *graph.Graph) {
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        for _, edge := range graphObj.Edges[node] {
            neighbour := edge.GetDestination()
            fmt.Printf("Current Node Value: %d, edge Destination Value: %d\n", node.GetValue(), neighbour.GetValue())
            fmt.Printf("Source Color: %s Neighbour Color: %s\n", node.GetColor(), neighbour.GetColor())
            if neighbour.GetColor() == "black" {
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
    queue = make([]*graph.Vertex, 0)
    initGraph := constructSampleGraph()
    graphVertices := initGraph.Vertices
    for range graphVertices {
        distance = append(distance, -1)
        predecessor = append(predecessor, -1)
    }
    startingNode := initGraph.GetVertexWithValue(0)
    distance[startingNode.GetValue()] = 0
    startingNode.SetColor("gray")
    queue = append(queue, startingNode)
    bfs_visit(initGraph)

    fmt.Printf("Queue: %v\n Distance: %v\n Prececessors: %v\n", queue, distance, predecessor)
}

    
    
        
        
    
