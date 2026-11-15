package main

import (
	"fmt"

	"github.com/Yahar4/internal/adjacency"
	"github.com/Yahar4/internal/csr"
	"github.com/Yahar4/internal/graphio"
)

var ident = func(i int) int { return i }

func main() {
	fmt.Println("=== 1. Построение CSR ===")

	// неориентированный граф
	c := csr.New[int](false)

	c.AddEdge(0, 1)
	c.AddEdge(1, 2)
	c.AddEdge(2, 3)

	c.AddVertex(4)

	fmt.Println("Вершины:      ", c.Vertices())
	fmt.Println("Соседи 1:     ", c.Neighbours(1))
	fmt.Println("Соседи 4:     ", c.Neighbours(4))

	fmt.Println("CSR - файл - CSR")

	csrPath := "./data/graph.csr"
	if err := graphio.SaveCSRFile(c, csrPath); err != nil {
		panic(err)
	}

	c2, err := graphio.LoadCSRFile[int](csrPath, ident)
	if err != nil {
		panic(err)
	}
	fmt.Println("Загружено. Вершины:", c2.Vertices())
	fmt.Println("Соседи 1 из файла:", c2.Neighbours(1))
	fmt.Println("Соседи 4 из файла:", c2.Neighbours(4))

	fmt.Println("Adjacency - файл - Adjacency")
	a := adjacency.New[int](false)
	a.AddEdge(0, 1)
	a.AddEdge(1, 2)
	a.AddEdge(2, 3)
	a.AddVertex(4)

	adjPath := "./data/graph.adj"
	if err := graphio.SaveAdjList(a, adjPath); err != nil {
		panic(err)
	}
	fmt.Println("Сохранено в", adjPath)

	a2, err := graphio.LoadAdjList[int](adjPath, ident)
	if err != nil {
		panic(err)
	}
	fmt.Println("Загружено. Вершины:", a2.Vertices())
	for _, v := range a2.Vertices() {
		fmt.Printf("  %d -> %v\n", v, a2.Neighbours(v))
	}

	fmt.Println("CSR - Adjacency")
	a3 := adjacency.New[int](false)
	for _, v := range c.Vertices() {
		a3.AddVertex(v)
	}
	for _, v := range c.Vertices() {
		for _, u := range c.Neighbours(v) {
			a3.AddEdge(v, u)
		}
	}
	for _, v := range a3.Vertices() {
		fmt.Printf("  %d -> %v\n", v, a3.Neighbours(v))
	}
}
