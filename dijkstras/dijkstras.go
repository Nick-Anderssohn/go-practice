package dijkstras

import (
	"fmt"
	"math"
	"practice/graph"

	"github.com/Nick-Anderssohn/go-ds/heap"
)

type Path struct {
	Edges       []*graph.Edge
	TotalWeight int
}

func FindShortestPath(g *graph.Graph, initialVertex *graph.Vertex, finalDestination *graph.Vertex) (*Path, error) {
	worker := initWorker(g, initialVertex, finalDestination)
	return worker.findShortestPath()
}

type worker struct {
	unvisitedVertices heap.Heap[*vertexWrapper]
	cur               *vertexWrapper
	finalDestination  *vertexWrapper

	vertexToWrapper map[*graph.Vertex]*vertexWrapper
}

type vertexWrapper struct {
	vertex *graph.Vertex
	path   *Path
}

func (w *worker) newVertexWrapper(vertex *graph.Vertex) *vertexWrapper {
	vw := &vertexWrapper{
		vertex: vertex,
		path: &Path{
			TotalWeight: math.MaxInt,
		},
	}

	if w.vertexToWrapper == nil {
		w.vertexToWrapper = map[*graph.Vertex]*vertexWrapper{}
	}
	w.vertexToWrapper[vertex] = vw

	return vw
}

func (v *vertexWrapper) LessThan(other any) bool {
	return v.path.TotalWeight < other.(*vertexWrapper).path.TotalWeight
}

func initWorker(g *graph.Graph, initialVertex *graph.Vertex, finalDestination *graph.Vertex) worker {
	unvisitedVertices, _ := heap.CreateHeap(heap.HeapTypeMin, []*vertexWrapper{})

	w := worker{
		unvisitedVertices: unvisitedVertices,
	}

	for v := range g.Vertices {
		wrapper := w.newVertexWrapper(v)

		if v == initialVertex {
			wrapper.path.TotalWeight = 0
			w.cur = wrapper
		} else {
			unvisitedVertices.Push(wrapper)
		}

		if v == finalDestination {
			w.finalDestination = wrapper
		}
	}

	return w
}

func (w *worker) findShortestPath() (*Path, error) {
	for w.unvisitedVertices.Len() > 0 {
		// First, update the tentative distances if a shorter path has been found to any neighbor
		for edge := range w.cur.vertex.Edges {
			neighbor := w.vertexToWrapper[edge.Vertex1]
			if neighbor == w.cur {
				neighbor = w.vertexToWrapper[edge.Vertex2]
			}

			potentialNewDist := w.cur.path.TotalWeight + edge.Weight

			if potentialNewDist < neighbor.path.TotalWeight {
				neighbor.path = &Path{
					TotalWeight: potentialNewDist,
					Edges:       append(w.cur.path.Edges, edge),
				}
			}
		}

		w.cur = w.unvisitedVertices.Pop()

		if w.cur == w.finalDestination {
			// we have found our path
			return w.cur.path, nil
		}
	}

	return nil, fmt.Errorf("no path found")
}
