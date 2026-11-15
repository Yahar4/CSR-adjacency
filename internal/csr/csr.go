package csr

import "fmt"

type CSRGraph[T comparable] struct {
	rowPtr   []int
	colIdx   []int
	toID     map[T]int
	fromID   []T
	directed bool
}

func New[T comparable](directed bool) *CSRGraph[T] {
	return &CSRGraph[T]{
		rowPtr:   []int{0},
		colIdx:   []int{},
		toID:     make(map[T]int),
		fromID:   []T{},
		directed: directed,
	}
}

func (g *CSRGraph[T]) AddVertex(v T) {
	if _, exists := g.toID[v]; exists {
		return
	}

	id := len(g.fromID)
	g.toID[v] = id
	g.fromID = append(g.fromID, v)

	g.rowPtr = append(g.rowPtr, g.rowPtr[id])
}

func (g *CSRGraph[T]) AddEdge(u, v T) {
	g.AddVertex(u)
	g.AddVertex(v)

	idU := g.toID[u]
	idV := g.toID[v]

	g.insertEdge(idU, idV)

	if !g.directed {
		g.insertEdge(idV, idU)
	}
}

func (g *CSRGraph[T]) Neighbours(v T) []T {
	id, ok := g.toID[v]
	if !ok {
		return nil
	}

	start := g.rowPtr[id]
	end := g.rowPtr[id+1]

	result := make([]T, end-start)
	for i := start; i < end; i++ {
		result[i-start] = g.fromID[g.colIdx[i]]
	}

	return result
}

func (g *CSRGraph[T]) Print() {
	fmt.Println("CSR Graph:")

	fmt.Printf("fromID = %v\n", g.fromID)
	fmt.Printf("rowPtr = %v\n", g.rowPtr)
	fmt.Printf("colIdx = %v\n", g.colIdx)

	for i, v := range g.fromID {
		start := g.rowPtr[i]
		end := g.rowPtr[i+1]

		fmt.Printf("%v -> ", v)

		for j := start; j < end; j++ {
			fmt.Printf("%v ", g.fromID[g.colIdx[j]])
		}

		fmt.Println()
	}
}

func (g *CSRGraph[T]) insertEdge(from, to int) {
	pos := g.rowPtr[from+1]

	g.colIdx = append(g.colIdx, 0)

	copy(g.colIdx[pos+1:], g.colIdx[pos:])

	g.colIdx[pos] = to

	for i := from + 1; i < len(g.rowPtr); i++ {
		g.rowPtr[i]++
	}
}

func (g *CSRGraph[T]) Vertices() []T {
	out := make([]T, len(g.fromID))
	copy(out, g.fromID)

	return out
}

func (g *CSRGraph[T]) Directed() bool { return g.directed }
