package graphio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Yahar4/internal/adjacency"
	"github.com/Yahar4/internal/csr"
)

func SaveAdjList[T comparable](g *adjacency.AdjacentGraph[T], path string) error {
	vs := g.Vertices()
	idx := map[T]int{}
	for i, v := range vs {
		idx[v] = i
	}

	var b strings.Builder
	fmt.Fprintln(&b, len(vs))
	for _, v := range vs {
		fmt.Fprintf(&b, "%d:", idx[v])
		for _, u := range g.Neighbours(v) {
			fmt.Fprintf(&b, " %d", idx[u])
		}
		fmt.Fprintln(&b)
	}
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func LoadAdjList[T comparable](path string, vertex func(int) T) (*adjacency.AdjacentGraph[T], error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	if !sc.Scan() {
		return nil, fmt.Errorf("пустой файл")
	}
	n, err := strconv.Atoi(strings.TrimSpace(sc.Text()))
	if err != nil {
		return nil, fmt.Errorf("n: %w", err)
	}

	g := adjacency.New[T](false)
	for i := 0; i < n; i++ {
		g.AddVertex(vertex(i))
	}

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		colon := strings.IndexByte(line, ':')
		if colon < 0 {
			return nil, fmt.Errorf("нет ':' в строке %q", line)
		}
		from, err := strconv.Atoi(line[:colon])
		if err != nil {
			return nil, fmt.Errorf("номер вершины %q: %w", line[:colon], err)
		}
		// соседи через пробел или запятую
		sep := func(r rune) bool { return r == ' ' || r == ',' }
		for _, s := range strings.FieldsFunc(line[colon+1:], sep) {
			to, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("сосед %q: %w", s, err)
			}
			g.AddEdge(vertex(from), vertex(to))
		}
	}
	return g, sc.Err()
}

func SaveCSRFile[T comparable](g *csr.CSRGraph[T], path string) error {
	vs := g.Vertices()
	idx := map[T]int{}
	for i, v := range vs {
		idx[v] = i
	}
	n := len(vs)

	row := make([]int, n+1)
	col := make([]int, 0)
	for i, v := range vs {
		row[i] = len(col)
		for _, u := range g.Neighbours(v) {
			col = append(col, idx[u])
		}
	}
	row[n] = len(col)

	var b strings.Builder
	fmt.Fprintln(&b, n, len(col))
	for _, x := range row {
		fmt.Fprintf(&b, "%d ", x)
	}
	fmt.Fprintln(&b)
	for _, x := range col {
		fmt.Fprintf(&b, "%d ", x)
	}
	fmt.Fprintln(&b)
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func LoadCSRFile[T comparable](path string, vertex func(int) T) (*csr.CSRGraph[T], error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var n, m int
	if _, err := fmt.Fscan(f, &n, &m); err != nil {
		return nil, fmt.Errorf("n m: %w", err)
	}
	row := make([]int, n+1)
	for i := range row {
		if _, err := fmt.Fscan(f, &row[i]); err != nil {
			return nil, fmt.Errorf("row_ptr[%d]: %w", i, err)
		}
	}
	col := make([]int, m)
	for i := range col {
		if _, err := fmt.Fscan(f, &col[i]); err != nil {
			return nil, fmt.Errorf("col_ind[%d]: %w", i, err)
		}
	}
	if row[n] != m {
		return nil, fmt.Errorf("row_ptr[n]=%d, но m=%d", row[n], m)
	}

	g := csr.New[T](false)
	for i := 0; i < n; i++ {
		g.AddVertex(vertex(i))
	}

	seen := map[[2]int]bool{}
	for i := 0; i < n; i++ {
		for j := row[i]; j < row[i+1]; j++ {
			a, b := i, col[j]
			if seen[[2]int{a, b}] {
				continue
			}
			seen[[2]int{a, b}] = true
			seen[[2]int{b, a}] = true
			g.AddEdge(vertex(a), vertex(b))
		}
	}
	return g, nil
}
