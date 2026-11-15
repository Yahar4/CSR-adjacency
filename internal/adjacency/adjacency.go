package adjacency

import "fmt"

type AdjacentGraph[T comparable] struct {
	vertices []T
	adj      map[T][]T // храним форматом тип -> массив того же типа
	directed bool
}

func New[T comparable](directed bool) *AdjacentGraph[T] {
	return &AdjacentGraph[T]{
		vertices: []T{},
		adj:      make(map[T][]T),
		directed: directed,
	}
}

func (g *AdjacentGraph[T]) AddVertex(v T) {
	if _, exists := g.adj[v]; !exists {
		g.adj[v] = []T{}
		g.vertices = append(g.vertices, v)
	}
}

// если граф направленный ставит значение в порядке в котором передаются значения
// в ином случае происходит наоборот
func (g *AdjacentGraph[T]) AddEdge(u, v T) {
	g.AddVertex(u)
	g.AddVertex(v)

	g.adj[u] = append(g.adj[u], v)
	if !g.directed {
		g.adj[v] = append(g.adj[v], u)
	}
}

// список смежных вершин
func (g *AdjacentGraph[T]) Neighbours(v T) []T {
	return g.adj[v]
}

func (g *AdjacentGraph[T]) Print() {
	for vertex, neighbours := range g.adj {
		fmt.Printf("%v -> %v \n", vertex, neighbours)
	}
}

func (g *AdjacentGraph[T]) Vertices() []T {
	out := make([]T, len(g.vertices))
	copy(out, g.vertices)

	return out
}

func (g *AdjacentGraph[T]) Directed() bool { return g.directed }
