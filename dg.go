package fabric

type Signal int

const (
	Waiting Signal = iota
	Started
	Complete
	Abort
	AbortRetry
	// TODO: PartialAbort (??)
)

// Dependency Graph Node
// every DGNode has an id, a state, and a set of Access Procedures
type DGNode interface {
	ID() int
	State() Signal
	ListProcedures() []ProceduresList
}

// Graph can be either UI DDAG, Temporal DAG or VDG
type Graph struct {
	Nodes []DGNode
	Edges map[DGNode][]DGNode // each node (id) has a list of node ids that it points too
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	var nodes []DGNode
	return &Graph{
		Nodes: nodes,
		Edges: make(map[DGNode][]DGNode),
	}
}

// GenerateGraph will create a graph given a list of nodes and map of edges
func GenerateGraph(nodes []DGNode, edges map[DGNode][]DGNode) *Graph {
	// TODO: should we just be supplied a list of node ids,
	// 		 and then generate nodes in a Waiting State (?)
	return &Graph{
		Nodes: nodes,
		Edges: edges,
	}
}

// CycleDetect will check whether a graph has cycles or not
func (g *Graph) CycleDetect() bool {
	var seen []DGNode
	var done []DGNode

	for _, v := range g.Nodes {
		if !contains(done, v) {
			result, _ := g.cycleDfs(v, seen, done)
			if result {
				return true
			}
		}
	}
	return false
}

// GetAdjacents will return the list of nodes a supplied node points too
func (g *Graph) GetAdjacents(node DGNode) []DGNode {
	return g.Edges[node]
}

// Recursive Depth-First-Search; used for Cycle Detection
func (g *Graph) cycleDfs(start DGNode, seen, done []DGNode) (bool, []DGNode) {
	seen = append(seen, start)
	adj := g.Edges[start]
	for _, v := range adj {
		if contains(done, v) {
			continue
		}

		if contains(seen, v) {
			return true, done
		}

		if result, done := g.cycleDfs(v, seen, done); result {
			return true, done
		}
	}
	seen = seen[:len(seen)-1]
	done = append(done, start)
	return false, done
}

// NOTE: this method cannot be used on UI DDAGs
func (g *Graph) RemoveNode(node DGNode) {
	// TODO:
	// 		remove node from node list
	// 		remove all of nodes edges from edges map
}

// Does append node even need an argument?
// should this just generate a node id and add it?
func (g *Graph) AppendNode(node DGNode) {
	// TODO: add node to nodes list
	// 		verify that node id is not already in list
	// 		verify that CDS sub-graph is not the same as in other UIs
}

// AppendEdge adds an edge that points from dependent to dependency
func (g *Graph) AppendEdge(source, dest DGNode) {
	// TODO: add an edge to edges map which is source dependent on destination
}

func (g *Graph) RemoveEdge(source, dest DGNode) {
	// TODO: removes the edge between source and dest in the edges map
}

func (g *Graph) AppendSubGraph(graph Graph) {
	// TODO:
	// 		add all nodes in new graph to existing graph
	// 		add all edges in new graph to existing graph
	// 		cycle detection on new total graph.
}

// NOTE: this method cannot be used on UI DDAGs
func (g *Graph) RemoveSubGraph(nodes []DGNode) {
	// TODO: removes list of nodes and all their edges
	//		(that they are sources for).
}

// UI Uniqueness Verification;
// slow-running processs, should only be called once
// when creating the UI dependency graph; can be called
// with the creation of each UI if needed for more "real-time"
// verification.
func (g *Graph) UniquenessVerification() bool {
	// TODO: verify that all UIs in the UI dependency
	// 		 graph are 'totality-unique'.
	return false
}

func (g *Graph) Coverage() bool {
	// TODO: returns whether or not every node and edge is addressed
	//		by at least one UI in our UI ddag.
	return true
}
