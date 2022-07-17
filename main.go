package main

type edges struct {
	src  int32
	dest int32
	tree map[int32][]int32
}

func (edges *edges) AddEdges() {
	if data, ok := edges.tree[edges.src]; ok {
		edges.tree[edges.src] = append(data, edges.dest)
	} else {
		edges.tree[edges.src] = []int32{edges.dest}
	}
}

func (edges *edges) doDFS(vertex int32) []int32 {
	visited := make([]int32, 0)
	return edges.UtilizeDFS(vertex, visited)
}

func (edges *edges) UtilizeDFS(vertex int32, visited []int32) []int32 {
	visited = append(visited, vertex)
	for _, value := range edges.tree[vertex] {
		if !Isitvisited(value, visited) {
			visited = edges.UtilizeDFS(value, visited)
		}
	}
	return visited
}
func Isitvisited(vertex int32, visited []int32) bool {
	for i := range visited {
		if visited[i] == vertex {
			return true
		}
	}
	return false
}

func (edges *edges) doBFS(vertex int32) []int32 {
	visited := make([]int32, 0)
	return edges.UtilizeBFS(vertex, visited)
}

func (edges *edges) UtilizeBFS(vertex int32, visited []int32) []int32 {
	queue := []int32{vertex}
	for len(queue) > 0 {
		currentvertex := queue[0]
		queue = queue[1:]
		if Isitvisited(currentvertex, visited) {
			continue
		}
		visited = append(visited, currentvertex)
		if len(edges.tree[currentvertex]) > 0 {
			queue = append(queue, edges.tree[currentvertex]...)
		}
	}
	return visited
}
