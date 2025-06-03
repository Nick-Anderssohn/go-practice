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

func FindShortestPath(initialVertex *graph.Vertex, finalDestination *graph.Vertex) (Path, error) {
	worker := initWorker(initialVertex, finalDestination)
	return worker.findShortestPath()
}

type worker struct {
	verticesVisited  map[*graph.Vertex]bool
	minHeap          heap.Heap[vertexWrapper]
	cur              vertexWrapper
	finalDestination *graph.Vertex

	vertexToWrapper map[*graph.Vertex]vertexWrapper
}

type vertexWrapper struct {
	vertex *graph.Vertex
	path   Path
}

func (w *worker) newVertexWrapper(vertex *graph.Vertex) vertexWrapper {
	vw := vertexWrapper{
		vertex: vertex,
		path: Path{
			TotalWeight: math.MaxInt,
		},
	}

	if w.vertexToWrapper == nil {
		w.vertexToWrapper = map[*graph.Vertex]vertexWrapper{}
	}
	w.vertexToWrapper[vertex] = vw

	return vw
}

func (v vertexWrapper) LessThan(other any) bool {
	return v.path.TotalWeight < other.(vertexWrapper).path.TotalWeight
}

func initWorker(initialVertex *graph.Vertex, finalDestination *graph.Vertex) worker {
	minHeap, _ := heap.CreateHeap(heap.HeapTypeMin, []vertexWrapper{})

	w := worker{
		verticesVisited:  map[*graph.Vertex]bool{},
		minHeap:          minHeap,
		finalDestination: finalDestination,
	}

	w.cur = w.newVertexWrapper(initialVertex)
	w.cur.path.TotalWeight = 0

	return w
}

func (w *worker) findShortestPath() (Path, error) {
	for w.cur.vertex != w.finalDestination {
		w.verticesVisited[w.cur.vertex] = true

		for edge := range w.cur.vertex.Edges {
			neighborVertex := edge.Vertex1
			if neighborVertex == w.cur.vertex {
				neighborVertex = edge.Vertex2
			}

			neighbor, wrapperExists := w.vertexToWrapper[neighborVertex]
			if !wrapperExists {
				neighbor = w.newVertexWrapper(neighborVertex)
			}

			potentialNewDist := w.cur.path.TotalWeight + edge.Weight

			if potentialNewDist < neighbor.path.TotalWeight {
				neighbor.path = Path{
					TotalWeight: potentialNewDist,
					Edges:       append(w.cur.path.Edges, edge),
				}

				// Prioritize the neighbor based off of its new distance
				w.minHeap.Push(neighbor)

				// update our saved wrapper.
				w.vertexToWrapper[neighbor.vertex] = neighbor
			}
		}

		// ensure we only visit each node once.
		for keepPopping := true; keepPopping; {
			if w.minHeap.Len() == 0 {
				return Path{}, fmt.Errorf("no path found")
			}

			w.cur = w.minHeap.Pop()

			if !w.verticesVisited[w.cur.vertex] {
				keepPopping = false
			} else if w.minHeap.Len() == 0 {
				return Path{}, fmt.Errorf("could not find a path to %s", w.finalDestination.ID)
			}
		}
	}

	return w.cur.path, nil
}
