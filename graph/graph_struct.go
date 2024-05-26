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
    Edges map[Vertex][]*Edge
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

func (e *Edge) GetDestination() *Vertex {
    return e.destination
}

func (e *Edge) GetWeight() int {
    return e.weight
}

func InitGraph() *Graph {
    return &Graph{
        Vertices: make([]*Vertex, 0),
        Edges: make(map[Vertex][]*Edge),
        VertexValuesMap: make(map[int]*Vertex),
    }
}

func GetVertex(value int, color string) *Vertex {
    return &Vertex{
        color: color,
        value: value,
    }
}

func GetEdge(destination *Vertex, weight int) *Edge {
    return &Edge{
        destination: destination,
        weight: weight,
    }
}

func (g *Graph) AddVertex(v GraphVertex) (bool, error) {
    vertexValue := v.GetValue()
    _, exists := g.VertexValuesMap[vertexValue]
    if exists {
        errMsg := fmt.Sprintf("%d already present as graph vertex. Returning", vertexValue)
        return false, errors.New(errMsg)
    }
    newVertex := &Vertex{
        value: vertexValue,
        color: "black",
    }
    g.Vertices = append(g.Vertices, newVertex)
    g.VertexValuesMap[vertexValue] = newVertex
    return true, nil
}

func (g *Graph) AddEdge (source *Vertex, destination *Vertex, weight int) (bool, error) {
    currentEdges, exists := g.Edges[*source]
    newEdge := &Edge{
        destination: destination,
        weight: weight,
    }
    if !exists {
        g.Edges[*source] = make([]*Edge, 0)
        g.Edges[*source] = append(g.Edges[*source], newEdge)
        return true, nil
    }
    
    if exists {
        for _, edge := range currentEdges {
            if destination == edge.destination {
                errMsg := fmt.Sprintf("%d -> %d already present", source.GetValue(), destination.GetValue())
                return false, errors.New(errMsg)
            }
        }
    }
    
    g.Edges[*source] = append(g.Edges[*source], newEdge)
    return true, nil
}

func (g *Graph) Print() {
    for key, vertexObj := range g.VertexValuesMap {
        fmt.Printf("%d: ", key)
        edges := g.Edges[*vertexObj]
        for _, edge := range edges {
            destinationValue := edge.GetDestination().GetValue()
            destinationWeight := edge.GetWeight()
            fmt.Printf("(%d, %d) ", destinationValue, destinationWeight)
        }
        fmt.Printf("\n")
    }
}
            

func ConstructGraphFromMap(vertexEdgeWeightMap map[int][]map[int]int) *Graph {
    newGraph := InitGraph()
    for value, edges := range vertexEdgeWeightMap {
        vertex := GetVertex(value, "black")
        _, vertexError := newGraph.AddVertex(vertex)
        if vertexError != nil {
            fmt.Printf("%s\n", vertexError)
        }
        for _, edge := range edges {
            for destinationValue, weight := range edge {
                destinationVertex := GetVertex(destinationValue, "black")
                destinationVertexCreated, destinationVertexError := newGraph.AddVertex(destinationVertex)
                fmt.Printf("%d Vertex Created: %v, Vertex Creation Error: %v\n", destinationValue, destinationVertexCreated, destinationVertexError)
                edgeCreated, edgeError := newGraph.AddEdge(vertex, destinationVertex, weight)
                fmt.Printf("%d -> %d Edge Created: %v, Edge Creation Error: %v\n", vertex.value, destinationVertex.value, edgeCreated, edgeError)
            }
        }
    }
    return newGraph
}
