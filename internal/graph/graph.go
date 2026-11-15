package graph

type Graph[T comparable] interface {
	AddVertex(v T)
	AddEdge(u, v T)
	Neighbours(v T) []T
	Print()
	Directed() bool
	Vertices() []T
}
