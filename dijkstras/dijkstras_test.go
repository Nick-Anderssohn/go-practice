package dijkstras

import (
	"slices"
	"testing"
)

var (
	VertexA = Vertex{"A"}
	VertexB = Vertex{"B"}
	VertexC = Vertex{"C"}
	VertexD = Vertex{"D"}
	VertexE = Vertex{"E"}
	VertexF = Vertex{"F"}
	VertexG = Vertex{"G"}
)

// This is the graph from this video: https://youtu.be/gdmfOwyQlcI?si=9IsOED253YMADyWO&t=250
var exampleGraph = Graph{
	InitialVertex:    &VertexA,
	FinalDestination: &VertexF,
	Edges: []*Edge{
		// Vertex A edges
		{
			Weight:  4,
			Vertex1: &VertexA,
			Vertex2: &VertexB,
		},
		{
			Weight:  3,
			Vertex1: &VertexA,
			Vertex2: &VertexC,
		},
		{
			Weight:  7,
			Vertex1: &VertexA,
			Vertex2: &VertexE,
		},

		// Vertex B edges
		// already have A-B
		{
			Weight:  6,
			Vertex1: &VertexB,
			Vertex2: &VertexC,
		},
		{
			Weight:  5,
			Vertex1: &VertexB,
			Vertex2: &VertexD,
		},

		// Vertex C edges
		// already have A-C
		// already have B-C
		{
			Weight:  11,
			Vertex1: &VertexC,
			Vertex2: &VertexD,
		},
		{
			Weight:  8,
			Vertex1: &VertexC,
			Vertex2: &VertexE,
		},

		// Vertex D edges
		// Already have B-D
		// Already have C-D
		{
			Weight:  2,
			Vertex1: &VertexD,
			Vertex2: &VertexE,
		},
		{
			Weight:  2,
			Vertex1: &VertexD,
			Vertex2: &VertexF,
		},
		{
			Weight:  10,
			Vertex1: &VertexD,
			Vertex2: &VertexG,
		},

		// Vertex E edges
		// Already have A-E
		// Already have C-E
		// Already have D-E
		{
			Weight:  5,
			Vertex1: &VertexE,
			Vertex2: &VertexG,
		},

		// Vertex F edges
		// Already have D-F
		{
			Weight:  3,
			Vertex1: &VertexF,
			Vertex2: &VertexG,
		},
	},
}

func TestDijkstras(t *testing.T) {
	result, err := FindShortestPath(&exampleGraph)

	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	expected := Path{
		TotalWeight: 11,
		Edges: []*Edge{
			exampleGraph.Edges[0],
			exampleGraph.Edges[4],
			exampleGraph.Edges[8],
		},
	}

	if result.TotalWeight != expected.TotalWeight || !slices.Equal(result.Edges, expected.Edges) {
		t.Fatalf("result is not equivalent to expected")
	}
}
