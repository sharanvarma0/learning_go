package main

import "learning_go/graph"

func main() {
    newGraphMap := map[int][]map[int]int{
        1: {
            map[int]int{2: 0}, map[int]int{3:1},
        },
        2: {
            map[int]int{3: 0},
        },
        3: {
            map[int]int{1: 1},
        },
    }
    graph := graph.ConstructGraphFromMap(newGraphMap)
    graph.Print()
}



