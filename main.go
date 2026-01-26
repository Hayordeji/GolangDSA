package main

import "fmt"

func main() {

	g := NewGraph()

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)

	fmt.Print("BFS: ")
	g.BFS(0)

	fmt.Print("\nDFS: ")
	g.DFS(0)
}

type Graph struct {
	adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adj: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Process node
		fmt.Print(node, " ")

		for _, neighbor := range g.adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	g.dfsHelper(start, visited)
}

func (g *Graph) dfsHelper(node int, visited map[int]bool) {
	visited[node] = true

	// Process node
	fmt.Print(node, " ")

	for _, neighbor := range g.adj[node] {
		if !visited[neighbor] {
			g.dfsHelper(neighbor, visited)
		}
	}
}
