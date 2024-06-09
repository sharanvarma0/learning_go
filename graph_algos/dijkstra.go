package main

import (
    "fmt"
    "learning_go/graph"
)

var visited []bool
var predecessor []int
var distance []int
var infinityValue int = -1
var nodeQueue []*graph.Vertex

func constructSampleGraph() *graph.Graph {
    graph_map := map[int][]map[int]int{
        0: {
            map[int]int{1: 3},
            map[int]int{6: 4},
            map[int]int{8: 5},
            },
        1: {
            map[int]int{2: 1},
            map[int]int{3: 1},
        },
        2: {
            map[int]int{10: 3},
            map[int]int{11: 6},
        },
        3: {
            map[int]int{4: 1},
            map[int]int{12: 1},
        },
        4: {
            map[int]int{5: 5},
            map[int]int{13: 5},
        },
        5: {
            map[int]int{6: 4},
            map[int]int{9: 4},
        },
        6: {
            map[int]int{7: 4},
        },
        7: {
            map[int]int{8: 3},
            map[int]int{9: 2},
        },
        8: {
            map[int]int{14: 2},
        },
        9: {
            map[int]int{15: 1},
        },
    }

    sampleGraph := graph.ConstructGraphFromMap(graph_map)
    sampleGraph.Print()
    return sampleGraph
}

func getNextUnvisitedMinNode(graph *graph.Graph) *graph.Vertex {
    minIndex := 0
    for _, vertex := range graph.Vertices {
        vertexValue := vertex.GetValue()
        if distance[vertexValue] < distance[minIndex] && !visited[vertexValue] {
            return vertex
        }
    }
    return nil
}
    

func initDijkstra(graph *graph.Graph, sourceNode *graph.Vertex) {
    visited = make([]bool, len(graph.Vertices))
    predecessor = make([]int, len(graph.Vertices))
    distance = make([]int, len(graph.Vertices))

    for i, _ := range visited {
        visited[i] = false
        predecessor[i] = -1
        distance[i] = infinityValue
    }
    
    distance[sourceNode.GetValue()] = 0
    nodeQueue = append(nodeQueue, sourceNode)
        
}

func dijkstra(graph *graph.Graph) { 
    for len(nodeQueue) > 0 {
        sourceNode := nodeQueue[0]
        nodeQueue = nodeQueue[1:]
        sourceValue := sourceNode.GetValue()
        for _, edge := range graph.Edges[sourceNode] {
            destinationVertexValue := edge.GetDestination().GetValue()
            newDistanceToDestination := distance[sourceValue] + edge.GetWeight()
            if distance[destinationVertexValue] == infinityValue || newDistanceToDestination < distance[destinationVertexValue] {
                distance[destinationVertexValue] = newDistanceToDestination
                predecessor[destinationVertexValue] = sourceNode.GetValue()
            }
        }
        visited[sourceNode.GetValue()] = true
        nextNode := getNextUnvisitedMinNode(graph)
        if nextNode != nil {
            nodeQueue = append(nodeQueue, nextNode)
        }
    }
}

func main() {
    sampleGraph := constructSampleGraph()
    nodeQueue = make([]*graph.Vertex, 0)
    sourceNode := sampleGraph.GetVertexWithValue(0)
    initDijkstra(sampleGraph, sourceNode)
    
    fmt.Printf("Source: %v\n", sourceNode)
    dijkstra(sampleGraph) 
    fmt.Printf("Predecessors: %v\n Distance: %v\n Visited: %v\n", predecessor, distance, visited)
}
