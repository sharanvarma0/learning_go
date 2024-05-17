package main

import (
    "fmt"
    "learning_go/priority_queue"
)

type LocalInt int

func (l LocalInt) GetPriority() int {
    return int(l)
}

func main() {
    priorityQueue := priority_queue.InitPriorityQueue()
    elements := []LocalInt{2,3,1,4,5,6,7,8}
    for _, elem := range elements {
        priorityQueue.Enqueue(elem)
    }
    
    elem, err := priorityQueue.Dequeue()
    for err == nil {
        fmt.Println(elem)
        elem, err = priorityQueue.Dequeue()
    }
}

    
    
