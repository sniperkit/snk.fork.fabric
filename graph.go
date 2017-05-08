package fabric

type Signal int

const (
	Incomplete Signal = iota
	Complete
	Abort
	AbortRetry
)

// Graph can be either UI DDAG, Temporal DAG or VDG
type Graph struct {
	Nodes []Node
	Edges map[int][]int // each node (id) has a list of node ids that it points too
}

type Node struct {
	Id    int
	State Signal
}

// TODO: differentiate between UI nodes, temporal nodes, and virtual nodes
// TODO: will need a UI uniqueness verification (for CDS node reference)
// TODO: how will we define a "CDS node"?

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	var nodes []Node
	return &Graph{
		Nodes: nodes,
		Edges: make(map[int][]int),
	}
}

// GenerateGraph will create a graph given a list of nodes and map of edges
func GenerateGraph(nodes []Node, edges map[int][]int) *Graph {
	return &Graph{
		Nodes: nodes,
		Edges: edges,
	}
}

// CycleDetect will check whether a graph has cycles or not
func (g *Graph) CycleDetect() bool {
	var seen []int
	var done []int

	for _, v := range g.Nodes {
		if !contains(done, v.Id) {
			result, done = g.cycleDfs(v.Id, seen, done)
			if result {
				return true
			}
		}
	}
	return false
}

// GetAdjacents will return the list of nodes a supplied node points too
func (g *Graph) GetAdjacents(node int) []int {
	return g.Edges[node]
}

// Recursive Depth-First-Search; used for Cycle Detection
func (g *Graph) cycleDfs(start int, seen, done []int) (bool, []int) {

	seen = append(seen, start)
	adj := g.Edges[start]
	for _, v := range adj {
		if contains(done, v) {
			continue
		}

		if contains(seen, i) {
			return true, done
		}

		if g.cycleDfs(v, seen, done) {
			return true, done
		}
	}
	seen = seen[:len(seen)-1]
	done = append(done, start)
	return false, done
}

// NOTE: this method cannot be used on UI DDAGs
func (g *Graph) RemoveNode(node int) {
	// TODO:
	// remove node from node list
	// remove all of nodes edges from edges map
}

// TODO: determine how Node IDs should be generated? Just incrementally?

// Does append node even need an argument?
// should this just generate a node id and add it?
func (g *Graph) AppendNode(node int) {
	// TODO: add node to nodes list
}

// AppendEdge adds an edge that points from dependent to dependency
func (g *Graph) AppendEdge(source, dest int) {
	// TODO: add an edge to edges map which is source dependent on destination
}

func (g *Graph) RemoveEdge(source, dest int) {
	// TODO: removes the edge between source and dest in the edges map
}

func (g *Graph) AppendSubGraph(graph Graph) {
	// TODO:
	// add all nodes in new graph to existing graph
	// add all edges in new graph to existing graph
	// cycle detection on new total graph
}

// NOTE: this method cannot be used on UI DDAGs
func (g *Graph) RemoveSubGraph(nodes []int) {
	// TODO: removes list of nodes and all their edges (that they are sources for)
}
