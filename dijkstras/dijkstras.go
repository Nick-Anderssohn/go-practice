package dijkstras

import (
	"fmt"
	"log"
	"math"

	"github.com/Nick-Anderssohn/go-ds/set"
)

type Vertex struct {
	ID string
}

type Edge struct {
	Weight  int
	Vertex1 *Vertex
	Vertex2 *Vertex
}

type Graph struct {
	Edges            []*Edge
	InitialVertex    *Vertex
	FinalDestination *Vertex
}

type worker struct {
	graph             *Graph
	edgeMap           map[*Vertex]set.Set[*Edge]
	paths             map[*Vertex]*Path
	visitedVertices   set.Set[*Vertex]
	unvisitedVertices set.Set[*Vertex]
	cur               *Vertex
}

type Path struct {
	Edges       []*Edge
	TotalWeight int
}

func FindShortestPath(graph *Graph) (Path, error) {
	worker := initWorker(graph)
	return worker.findShortestPath()
}

func initWorker(graph *Graph) worker {
	log.Printf("graph: %+v\n", *graph)
	edgeMap := map[*Vertex]set.Set[*Edge]{}
	paths := map[*Vertex]*Path{}
	visitedVertices := set.Set[*Vertex]{graph.InitialVertex: true}
	unvisitedVertices := set.Set[*Vertex]{}

	for _, edge := range graph.Edges {
		edgeMap[edge.Vertex1] = set.Put(edgeMap[edge.Vertex1], edge)
		edgeMap[edge.Vertex2] = set.Put(edgeMap[edge.Vertex2], edge)
		paths[edge.Vertex1] = &Path{TotalWeight: math.MaxInt}
		paths[edge.Vertex2] = &Path{TotalWeight: math.MaxInt}

		if edge.Vertex1 != graph.InitialVertex {
			set.Put(unvisitedVertices, edge.Vertex1)
		}

		if edge.Vertex2 != graph.InitialVertex {
			set.Put(unvisitedVertices, edge.Vertex2)
		}
	}

	// Initial one needs a weight of 0
	paths[graph.InitialVertex] = &Path{TotalWeight: 0}

	return worker{
		graph:             graph,
		edgeMap:           edgeMap,
		paths:             paths,
		visitedVertices:   visitedVertices,
		unvisitedVertices: unvisitedVertices,
		cur:               graph.InitialVertex,
	}
}

func (w *worker) findShortestPath() (Path, error) {
	log.Printf("searching for shortest path...")
	for len(w.unvisitedVertices) > 0 {
		curPath := w.paths[w.cur]

		log.Printf("cur vertex: %#v\n", *w.cur)
		log.Printf("cur path: %#v\n", *curPath)

		// First, update the tentative distances if a shorter path has been found to any neighbor
		for edge := range w.edgeMap[w.cur] {
			log.Printf("checking edge: %#v\n", *edge)
			neighbor := edge.Vertex1
			if neighbor == w.cur {
				neighbor = edge.Vertex2
			}

			existingDist := w.paths[neighbor].TotalWeight
			potentialNewDist := curPath.TotalWeight + edge.Weight

			if potentialNewDist < existingDist {
				w.paths[neighbor] = &Path{
					Edges:       append(curPath.Edges, edge),
					TotalWeight: potentialNewDist,
				}
			}
		}

		// mark cur as visited
		set.Put(w.visitedVertices, w.cur)
		set.Remove(w.unvisitedVertices, w.cur)

		if set.Exists(w.visitedVertices, w.graph.FinalDestination) {
			// we have found our path
			return *w.paths[w.graph.FinalDestination], nil
		}

		var next *Vertex
		for vertex := range w.unvisitedVertices {
			if next == nil || w.paths[vertex].TotalWeight < w.paths[next].TotalWeight {
				next = vertex
			}
		}

		if next == nil {
			return Path{}, fmt.Errorf("no path found")
		}

		w.cur = next
	}

	return Path{}, fmt.Errorf("no path found")
}
