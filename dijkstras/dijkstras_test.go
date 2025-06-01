package dijkstras

import (
	"practice/graph"
	"slices"
	"testing"

	"github.com/Nick-Anderssohn/go-ds/set"
)

var (
	VertexA = &graph.Vertex{ID: "A"}
	VertexB = &graph.Vertex{ID: "B"}
	VertexC = &graph.Vertex{ID: "C"}
	VertexD = &graph.Vertex{ID: "D"}
	VertexE = &graph.Vertex{ID: "E"}
	VertexF = &graph.Vertex{ID: "F"}
	VertexG = &graph.Vertex{ID: "G"}

	edges = []*graph.Edge{
		// Vertex A edges
		graph.InitEdge(4, VertexA, VertexB),
		graph.InitEdge(3, VertexA, VertexC),
		graph.InitEdge(7, VertexA, VertexE),

		// Vertex B edges
		// already have A-B
		graph.InitEdge(6, VertexB, VertexC),
		graph.InitEdge(5, VertexB, VertexD),

		// Vertex C edges
		// already have A-C
		// already have B-C
		graph.InitEdge(11, VertexC, VertexD),
		graph.InitEdge(8, VertexC, VertexE),

		// Vertex D edges
		// Already have B-D
		// Already have C-D
		graph.InitEdge(2, VertexD, VertexE),
		graph.InitEdge(2, VertexD, VertexF),
		graph.InitEdge(10, VertexD, VertexG),

		// Vertex E edges
		// Already have A-E
		// Already have C-E
		// Already have D-E
		graph.InitEdge(5, VertexE, VertexG),

		// Vertex F edges
		// Already have D-F
		graph.InitEdge(3, VertexF, VertexG),
	}

	expectedTotalWeight = 11

	expectedPath1 = []*graph.Edge{
		edges[0],
		edges[4],
		edges[8],
	}

	expectedPath2 = []*graph.Edge{
		edges[2],
		edges[7],
		edges[8],
	}
)

// This is the graph from this video: https://youtu.be/gdmfOwyQlcI?si=9IsOED253YMADyWO&t=250
var exampleGraph = graph.InitGraph(set.FromSlice(edges))

func TestDijkstras(t *testing.T) {
	result, err := FindShortestPath(&exampleGraph, VertexA, VertexF)

	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	if result.TotalWeight != expectedTotalWeight {
		t.Errorf("expected total weight of %d but got %d", expectedTotalWeight, result.TotalWeight)
	}

	// It is valid for the result to be either expectedPath1 or expectedPath2 since both paths
	// have a total weight of 11.
	if !slices.Equal(result.Edges, expectedPath1) && !slices.Equal(result.Edges, expectedPath2) {
		t.Fatalf("result is not equivalent to expected")
	}
}
