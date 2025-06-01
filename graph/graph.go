package graph

import "github.com/Nick-Anderssohn/go-ds/set"

type Vertex struct {
	ID string

	// Will be initialized/updated by InitEdge function
	Edges set.Set[*Edge]
}

type Edge struct {
	Weight  int
	Vertex1 *Vertex
	Vertex2 *Vertex
}

type Graph struct {
	Vertices set.Set[*Vertex]
	Edges    set.Set[*Edge]
}

func InitGraph(edges set.Set[*Edge]) Graph {
	vertices := set.Set[*Vertex]{}

	for edge := range edges {
		set.Put(vertices, edge.Vertex1)
		set.Put(vertices, edge.Vertex2)
	}

	return Graph{
		Vertices: vertices,
		Edges:    edges,
	}
}

func InitEdge(weight int, vertex1 *Vertex, vertex2 *Vertex) *Edge {
	edge := &Edge{
		Weight:  weight,
		Vertex1: vertex1,
		Vertex2: vertex2,
	}

	vertex1.Edges = set.Put(vertex1.Edges, edge)
	vertex2.Edges = set.Put(vertex2.Edges, edge)

	return edge
}
