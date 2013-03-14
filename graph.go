package graph
// In computer science, a graph is an abstract data type that is meant to 
// implement the graph and hypergraph concepts from mathematics.
type Interface interface {
    Adjacent(Node, Node) (bool)
    Add(Node, Node)
    Delete(Node, Node)
    Neighbors(Node) (map[Node]Edge)
}

// A graph data structure consists of a finite (and possibly mutable) set of 
// ordered pairs, called edges or arcs, of certain entities called nodes or 
// vertices. As in mathematics, an edge (x,y) is said to point or go from x to 
// y. The nodes may be part of the graph structure, or may be external entities 
// represented by integer indices or references.
type Node interface {
}

// A graph data structure may also associate to each edge some edge value, such 
// as a symbolic label or a numeric attribute (cost, capacity, length, etc.).
type Edge interface {
}

// ## Operations
// The basic operations provided by a graph data structure G usually include:
func New() (g Interface) {
    g = &adjacencyGraph{make(map[Node]map[Node]Edge)}
    return
}

// ### Adjacent
// Tests whether there is an edge from node x to node y.
func (g *adjacencyGraph) Adjacent(x Node, y Node) (result bool) {
    if edgeNodes, ok := g.storage[x]; ok {
        _, result = edgeNodes[y]
    }
    return
}

// ### Neighbors
// lists all nodes y such that there is an edge from x to y.
func (g *adjacencyGraph) Neighbors(x Node) (y map[Node]Edge){
    y = g.storage[x]
    return
}

// ### Add
// adds to G the edge from x to y, if it is not there.
func (g *adjacencyGraph) Add(x Node, y Node) {
    if edgeNodes, ok := g.storage[x]; ok {
        edgeNodes[y] = y
    } else {
        g.storage[x] = map[Node]Edge{y:nil}
    }
}

// ### Delete
// removes the edge from x to y, if it is there.
func (g *adjacencyGraph) Delete(x Node, y Node) {
    if edgeNodes, ok := g.storage[x]; ok {
        delete(edgeNodes, y)
    }
}

// ## Representations
// Different data structures for the representation of graphs are used in
// practice:

// ### Adjacency list
// Vertices are stored as records or objects, and every vertex stores a list of
// adjacent vertices. This data structure allows the storage of additional data
// on the vertices

type adjacencyGraph struct {
    storage map[Node]map[Node]Edge
}

// ### Incidence list 
// Vertices and edges are stored as records or objects. Each vertex stores its
// incident edges, and each edge stores its incident vertices. This data
// structure allows the storage of additional data on vertices and edges.

// ### Adjacency matrix 
// A two-dimensional matrix, in which the rows represent source vertices and
// columns represent destination vertices. Data on edges and vertices must be
// stored externally. Only the cost for one edge can be stored between each 
// pair of vertices.

// ### Incidence matrix 
// A two-dimensional Boolean matrix, in which the rows represent the vertices 
// and columns represent the edges. The entries indicate whether the vertex at 
// a row is incident to the edge at a column.

// ## Algorithms
//
// ### Dijkstra
func Dijkstra(g Interface, source Node) {
    // Assign to every node a tentative distance value: set it to zero for our 
    // initial node and to infinity for all other nodes.
    dist := map[Node]int { source: 0 }
    // Mark all nodes unvisited.
    visited := map[Node]Node{}
    // Set the initial node as current.
    for current := source; current != nil; {
        // For the current node, consider all of its unvisited neighbors and 
        // calculate their tentative distances.
        distance := dist[current]
        for node, _ := range g.Neighbors(current) {
            // Currently all edges have a distance of one.
            newdistance := distance + 1
            // If this distance is less than the previously recorded tentative 
            // distance of B, then overwrite that distance.
            if olddistance, ok := dist[node];!ok || newdistance < olddistance {
                dist[node] = newdistance
            }
        }
        // When we are done considering all of the neighbors of the current 
        // node, mark the current node as visited.
        visited[current] = current
        current = nil

        // Select the unvisited node that is marked with the smallest 
        // tentative distance, and set it as the current node.
        smallest := int(^uint(0) >> 1) 
        for node, distance := range dist {
            if _, ok := visited[node]; !ok && smallest < distance {
                smallest = distance
                current = node
            }
        }
    }
}
