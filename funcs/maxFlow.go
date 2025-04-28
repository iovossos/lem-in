package funcs

// This file implements a max flow solution to find vertex-disjoint paths
// using the standard technique of vertex splitting and the Edmonds–Karp algorithm.

// Edge represents a directed edge in the flow network.
type Edge struct {
	from     string
	to       string
	capacity int
	flow     int
	rev      *Edge
}

// addEdge adds a forward edge (with capacity) and a corresponding reverse edge (with 0 capacity)
// to the flow network.
func addEdge(graph map[string][]*Edge, from, to string, capacity int) {
	// Create forward and reverse edges.
	forward := &Edge{from: from, to: to, capacity: capacity, flow: 0}
	reverse := &Edge{from: to, to: from, capacity: 0, flow: 0}
	forward.rev = reverse
	reverse.rev = forward
	graph[from] = append(graph[from], forward)
	graph[to] = append(graph[to], reverse)
}

// transformIn returns the name of the "in" version of a vertex.
// For the start and end vertices, no splitting is performed.
func transformIn(node, start, end string) string {
	if node == start || node == end {
		return node
	}
	return node + "_in"
}

// transformOut returns the name of the "out" version of a vertex.
// For the start and end vertices, no splitting is performed.
func transformOut(node, start, end string) string {
	if node == start || node == end {
		return node
	}
	return node + "_out"
}

// BuildFlowNetwork creates the flow network graph using vertex splitting.
// It returns both the network and a nodeMap to translate transformed node names back to original names.
func BuildFlowNetwork(connections map[string][]string, start, end string) (map[string][]*Edge, map[string]string) {
	network := make(map[string][]*Edge)
	nodeMap := make(map[string]string) // Maps transformed node names to original names.

	// Gather all vertices from the connections.
	vertexSet := make(map[string]bool)
	for u, neighbors := range connections {
		vertexSet[u] = true
		for _, v := range neighbors {
			vertexSet[v] = true
		}
	}

	// For every vertex, if it is not start or end, split it into "in" and "out" nodes.
	for v := range vertexSet {
		if v == start || v == end {
			// For source and sink, no split is needed.
			if _, ok := network[v]; !ok {
				network[v] = []*Edge{}
			}
			nodeMap[v] = v
		} else {
			vin := transformIn(v, start, end)
			vout := transformOut(v, start, end)
			network[vin] = []*Edge{}
			network[vout] = []*Edge{}
			// Connect vin to vout with capacity 1 to enforce the vertex constraint.
			addEdge(network, vin, vout, 1)
			nodeMap[vin] = v
			nodeMap[vout] = v
		}
	}

	// To simulate undirected edges, use a set to add each edge only once.
	edgeAdded := make(map[string]bool)
	for u, neighbors := range connections {
		for _, v := range neighbors {
			// Create a key for the undirected edge based on sorted vertex names.
			key := ""
			if u < v {
				key = u + "_" + v
			} else {
				key = v + "_" + u
			}
			if !edgeAdded[key] {
				edgeAdded[key] = true
				// Determine the proper transformed names for u and v.
				uOut := u
				if u != start && u != end {
					uOut = transformOut(u, start, end)
				}
				vIn := v
				if v != start && v != end {
					vIn = transformIn(v, start, end)
				}
				// Add edge from u's output to v's input.
				addEdge(network, uOut, vIn, 1)

				// Do the reverse direction to simulate an undirected tunnel.
				vOut := v
				if v != start && v != end {
					vOut = transformOut(v, start, end)
				}
				uIn := u
				if u != start && u != end {
					uIn = transformIn(u, start, end)
				}
				addEdge(network, vOut, uIn, 1)
			}
		}
	}
	return network, nodeMap
}

// bfs finds an augmenting path in the flow network and fills in the parent mapping.
// Returns true if a path from source to sink is found.
func bfs(network map[string][]*Edge, source, sink string, parent map[string]*Edge) bool {
	// Clear the parent map.
	for key := range parent {
		delete(parent, key)
	}
	queue := []string{source}
	visited := make(map[string]bool)
	visited[source] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, edge := range network[u] {
			// Compute the residual capacity.
			if edge.capacity-edge.flow > 0 && !visited[edge.to] {
				visited[edge.to] = true
				parent[edge.to] = edge
				if edge.to == sink {
					return true
				}
				queue = append(queue, edge.to)
			}
		}
	}
	return false
}

// MaxFlow computes the maximum flow from source to sink using the Edmonds–Karp algorithm.
func MaxFlow(network map[string][]*Edge, source, sink string) int {
	parent := make(map[string]*Edge)
	maxFlow := 0

	// While there is an augmenting path, send flow (which is always 1 in this construction).
	for bfs(network, source, sink, parent) {
		flow := 1
		// Traverse the path from sink back to source and update flows.
		for v := sink; v != source; {
			edge := parent[v]
			edge.flow += flow
			edge.rev.flow -= flow
			v = edge.from
		}
		maxFlow += flow
	}
	return maxFlow
}

// dfsExtract uses depth-first search to extract one path from source to sink along edges with flow > 0.
// As a path is found, the flow on the used edges is decremented to avoid using them again.
func dfsExtract(network map[string][]*Edge, u, sink string, path *[]string) bool {
	// Append current node to path.
	*path = append(*path, u)
	if u == sink {
		return true
	}
	for _, edge := range network[u] {
		// Follow only edges that have flow remaining.
		if edge.flow > 0 {
			// Decrement the flow to mark this edge as used.
			edge.flow--
			if dfsExtract(network, edge.to, sink, path) {
				return true
			}
			// Backtrack if the path did not lead to sink.
		}
	}
	// Remove u from path if no valid continuation exists.
	*path = (*path)[:len(*path)-1]
	return false
}

// ExtractPaths retrieves all vertex-disjoint paths from the flow network.
// It uses the computed flows (each unit of flow corresponds to one path),
// and converts the transformed node names back to original room names.
func ExtractPaths(network map[string][]*Edge, nodeMap map[string]string, source, sink string, flow int) [][]string {
	var paths [][]string
	for i := 0; i < flow; i++ {
		var path []string
		if dfsExtract(network, source, sink, &path) {
			// Convert transformed node names to the original names.
			var originalPath []string
			for _, node := range path {
				if node == source || node == sink {
					originalPath = append(originalPath, node)
				} else {
					originalPath = append(originalPath, nodeMap[node])
				}
			}
			// Clean up consecutive duplicates that may result from the splitting.
			cleanedPath := []string{}
			for j, name := range originalPath {
				if j == 0 || name != originalPath[j-1] {
					cleanedPath = append(cleanedPath, name)
				}
			}
			paths = append(paths, cleanedPath)
		}
	}
	return paths
}

// VertexDisjointPaths computes the vertex-disjoint paths using the max flow approach.
// It returns a slice of paths (each a slice of room names) and the maximum flow value.
func VertexDisjointPaths(connections map[string][]string, start, end string) ([][]string, int) {
	// Build the flow network with vertex splitting.
	network, nodeMap := BuildFlowNetwork(connections, start, end)
	// Compute the maximum flow. The flow value equals the number of vertex-disjoint paths.
	maxFlowValue := MaxFlow(network, start, end)
	// Extract the actual paths from the flow network.
	paths := ExtractPaths(network, nodeMap, start, end, maxFlowValue)
	return paths, maxFlowValue
}
