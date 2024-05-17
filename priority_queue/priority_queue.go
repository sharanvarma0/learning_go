/*****************************
This package implements a simple PriorityQueue as a heap.
The Queue is itself a structure with a field for the heap slice.
The Elements in the slice should satisfy the PriorityQueueElement Interface

Example Usage:
    queue := priority_queue.InitPriorityQueue()
    queue.Enqueue(<element>)
    elem, err := queue.Dequeue()
    if err != nil {
        Handle
    }
    fmt.Printf("%v\n", elem)

This was written for a more elaborate use in the implementation of Dijkstra's Algorithm
as a sort of experiment
********************************/


package priority_queue

import "errors"

type PriorityQueueElement interface {
    GetPriority() int
}

type Queue struct {
    heap []PriorityQueueElement
}

func InitPriorityQueue() *Queue  {
    newHeap := make([]PriorityQueueElement, 0)
    newQueue := &Queue{
        heap: newHeap,
    }
    return newQueue
}
        

func Heapify(q *Queue, index int, length int) {
    left, right := 2*index + 1, 2*index + 2
    largest := index
    largestPriority := q.heap[largest].GetPriority()
    if (left < length) {
        leftPriority := q.heap[left].GetPriority() 
        if leftPriority > largestPriority {
            largest = left
            largestPriority = leftPriority
        }
    }
    if (right < length) {
        rightPriority := q.heap[right].GetPriority()
        if rightPriority > largestPriority {
            largestPriority = rightPriority
            largest = right
        }
    }

    if largest != index {
        q.heap[largest], q.heap[index] = q.heap[index], q.heap[largest]
        Heapify(q, largest, length)
    }
}

func RebuildHeap(q *Queue) {
    length := len(q.heap)
    
    for index := int(length/2) - 1; index >= 0; index-- {
        Heapify(q, index, length)
    }
}

func (q *Queue) Enqueue(element PriorityQueueElement) bool {
    q.heap = append(q.heap, element)
    RebuildHeap(q)
    return true
}

func (q *Queue) Dequeue() (PriorityQueueElement, error) {
    if len(q.heap) == 0 {
        return nil, errors.New("Empty Queue")
    }
    head := q.heap[0]
    q.heap = q.heap[1:]
    RebuildHeap(q)
    return head, nil
}
    
    
    
        
        
