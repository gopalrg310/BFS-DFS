# Depth First Search and Breadth First Search

This will be used to do BFS and DFS in simplest format.

# To do Depth First Search
Try doDFS with given map of tree

Depth First Traversal (or Search) for a graph is similar to Depth First Traversal of a tree. The only catch here is, unlike trees, graphs may contain cycles (a node may be visited twice). To avoid processing a node more than once, use a boolean visited array. 


# To do Breadth First Search
Try doBFS with given map of tree

About BFS Breadth-First Traversal (or Search) for a graph is similar to Breadth-First Traversal of a tree. The only catch here is
that, unlike trees, graphs may contain cycles, so we may come to the same node again. To avoid processing a node more than once, we use a boolean visited array. For simplicity, it is assumed that all vertices are reachable from the starting vertex. BFS uses a queue data structure for traversal.


When there is a array of map to 

Access it by `go get github.com/gopalrg310/BFS-DFS`

To Run test cases
`go test -v`