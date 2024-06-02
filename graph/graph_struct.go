package graph

import (
    "errors"
    "fmt"
)

type Vertex struct {
    value int
    color string
}

type Edge struct {
    destination *Vertex
    weight int
}

type Graph struct {
    Vertices []*Vertex
    Edges map[*Vertex][]*Edge
    VertexValuesMap map[int]*Vertex
}

type GraphVertex interface {
    GetValue() int
}

func (v *Vertex) GetValue() int {
    return v.value
}

func (v *Vertex) GetColor() string {
    return v.color
}

func (v *Vertex) SetColor(color string) bool {
    v.color = color
    return true
}

func (v *Vertex) SetValue(value int) bool {
    v.value = value
    return true
}

func (e *Edge) GetDestination() *Vertex {
    return e.destination
}

func (e *Edge) GetWeight() int {
    return e.weight
}

func InitGraph() *Graph {
    return &Graph{
        Vertices: make([]*Vertex, 0),
        Edges: make(map[*Vertex][]*Edge),
        VertexValuesMap: make(map[int]*Vertex),
    }
}

func CreateNewVertex(value int, color string) *Vertex {
    return &Vertex{
        color: color,
        value: value,
    }
}

func CreateNewEdge(destination *Vertex, weight int) *Edge {
    return &Edge{
        destination: destination,
        weight: weight,
    }
}


func (g *Graph) GetVertexWithValue(value int) *Vertex {
    vertexPtr, exists := g.VertexValuesMap[value]
    if !exists {
        fmt.Printf("%d Vertex does not exist in the graph. Please create it if required")
        return nil
    }
    return vertexPtr
}
    

func (g *Graph) AddVertexToGraph(v GraphVertex) (*Vertex, error) {
    vertexValue := v.GetValue()
    existingVertex, exists := g.VertexValuesMap[vertexValue]
    if exists {
        errMsg := fmt.Sprintf("%d already present as graph vertex. Returning", vertexValue)
        return existingVertex, errors.New(errMsg)
    }
    newVertex := CreateNewVertex(vertexValue, "white")
    g.Vertices = append(g.Vertices, newVertex)
    g.VertexValuesMap[vertexValue] = newVertex
    return newVertex, nil
}


func (g *Graph) checkForExistingEdge(source, destination *Vertex) (srcToDest bool, destToSrc bool) {
    for _, edge := range g.Edges[source] {
        if destination == edge.destination {
            srcToDest = true
        }
    }
    for _, edge := range g.Edges[destination] {
        if edge.destination == source {
            destToSrc = true
        }
    }

    return
}

func (g *Graph) AddEdgeToGraph(source *Vertex, destination *Vertex, weight int) (bool, error) {
    _, exists := g.Edges[source]
    _, destEdgesExists := g.Edges[destination]
    if !exists {
        g.Edges[source] = make([]*Edge, 0)
    }
    if !destEdgesExists {
        g.Edges[destination] = make([]*Edge, 0)
    }
   
    srcToDest, destToSrc := g.checkForExistingEdge(source, destination)

    if !srcToDest {
        newEdge := CreateNewEdge(destination, weight)
        g.Edges[source] = append(g.Edges[source], newEdge)
    }

    if weight == 0 {
        if !destToSrc {
            destinationEdge := CreateNewEdge(source, weight)
            g.Edges[destination] = append(g.Edges[destination], destinationEdge)
            fmt.Println("Weight Passed: 0. Added Bidirectional Edge")
        }
    }
    return true, nil
}

func (g *Graph) Print() {
    fmt.Printf("Vertices: %v\n Edges: %v\n", g.Vertices, g.Edges)
    for key, vertexObj := range g.VertexValuesMap {
        fmt.Printf("%d: ", key)
        edges := g.Edges[vertexObj]
        fmt.Printf("Edges: %v\n", edges)
        for _, edge := range edges {
            destinationValue := edge.GetDestination().GetValue()
            destinationWeight := edge.GetWeight()
            fmt.Printf("(%d, %d) ", destinationValue, destinationWeight)
        }
        fmt.Printf("\n")
    }
}
            

func ConstructGraphFromMap(vertexEdgeMap map[int][]map[int]int) *Graph {
    newGraph := InitGraph()
    for value, edges := range vertexEdgeMap {
        vertex := CreateNewVertex(value, "white")
        srcVertexPtr, vertexError := newGraph.AddVertexToGraph(vertex)
        if vertexError != nil {
            fmt.Printf("Vertex Creation Error: %s\n", vertexError)
        }
        for _, edge := range edges {
            for destinationValue, weight := range edge {
                destinationVertex := CreateNewVertex(destinationValue, "white")
                destinationVertexPtr, destinationVertexError := newGraph.AddVertexToGraph(destinationVertex)
                fmt.Printf("%d Vertex Created, Vertex Creation Error: %v\n", destinationValue, destinationVertexError)
                if destinationVertexPtr != nil {
                    edgeCreated, edgeError := newGraph.AddEdgeToGraph(srcVertexPtr, destinationVertexPtr, weight)
                    fmt.Printf("%d -> %d Edge Created: %v, Edge Creation Error: %v\n", vertex.value, destinationVertex.value, edgeCreated, edgeError)
                }
            }
        }
    }
    return newGraph
}
