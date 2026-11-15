package convert

import (
	"github.com/Yahar4/internal/adjacency"
	"github.com/Yahar4/internal/csr"
)

func CSRToAdjacency[T comparable](g *csr.CSRGraph[T]) *adjacency.AdjacentGraph[T] {
	out := adjacency.New[T](g.Directed())

	for _, v := range g.Vertices() {
		for _, u := range g.Neighbours(v) {
			out.AddEdge(v, u)
		}
	}

	return out
}

func AdjacencyToCSR[T comparable](g *adjacency.AdjacentGraph[T]) *csr.CSRGraph[T] {
	out := csr.New[T](g.Directed())

	for _, v := range g.Vertices() {
		out.AddVertex(v)
	}

	seen := make(map[[2]T]bool)
	for _, v := range g.Vertices() {
		for _, u := range g.Neighbours(v) {
			if !g.Directed() && seen[[2]T{u, v}] {
				continue
			}

			seen[[2]T{v, u}] = true
			out.AddEdge(v, u)
		}
	}

	return out
}
