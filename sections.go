package fabric

type Section interface {
	ListNodes() NodeList
	ListEdges() EdgeList
}

/*
	DS = Data Structure; used when a UI will have
	access to entire CDS.
*/

type DS struct {
	Nodes NodeList
	Edges EdgeList
}

// TODO: accept CDS as arguement and return entire CDS
//		as a section.
func NewDS(nodes NodeList, edges EdgeList) *DS {
	return &DS{
		Nodes: nodes,
		Edges: edges,
	}
}

func (s *DS) ListNodes() NodeList {
	return s.Nodes
}

func (s *DS) ListEdges() EdgeList {
	return s.Edges
}

/* Sub-graphs are non-disjoint collections of nodes and edges */
type Subgraph struct {
	Nodes NodeList
	Edges EdgeList
}

func NewSubgraph(nodes []Node) *Subgraph {

	// TODO: will grab all edges from nodes that connect to
	//		other nodes that are in our list.

	var edges EdgeList

	return &Subgraph{
		Nodes: nodes,
		Edges: edges,
	}
}

func (s *Subgraph) ListNodes() NodeList {
	return s.Nodes
}

func (s *Subgraph) ListEdges() EdgeList {
	return s.Edges
}

/*
	Branches are all nodes and edges for a particuliar branch
	(usually of a tree graph)
	A branch is technically a sub-graph as well.
*/
type Branch struct {
	Nodes NodeList
	Edges EdgeList
}

func NewBranch(root Node) *Branch {
	var nodes NodeList
	var edges EdgeList

	// TODO: grab all children nodes recursively
	return &Branch{
		Nodes: nodes,
		Edges: edges,
	}
}

func (b *Branch) ListNodes() NodeList {
	return b.Nodes
}

func (b *Branch) ListEdges() EdgeList {
	return b.Edges
}

/*
	Partitions are only for linear CDSs
	(i.e. each node can only have at most 2 edges)
*/
type Partition struct {
	Nodes NodeList
	Edges EdgeList
}

func NewPartition(start, end Node) *Partition {
	// TODO: adds all nodes between and including the start
	//		and end node; will also grab all edges for these
	//		nodes.
	var nodes NodeList
	var edges EdgeList

	return &Partition{
		Nodes: nodes,
		Edges: edges,
	}
}

func (p *Partition) ListNodes() NodeList {
	return p.Nodes
}

func (p *Partition) ListEdges() EdgeList {
	return p.Edges
}

/* Subsets are used for generic node selection (but not generic edge selection) */
type Subset struct {
	Nodes NodeList
	Edges EdgeList
}

func NewSubset(nodes NodeList) *Subset {
	// TODO: grab all (and only all) edges that are connected
	//		to a node in the list of nodes supplied.
	var edges EdgeList

	return &Subset{
		Nodes: nodes,
		Edges: edges,
	}
}

func (s *Subset) ListNodes() NodeList {
	return s.Nodes
}

func (s *Subset) ListEdges() EdgeList {
	return s.Edges
}

/* Disjoints are a collection of arbitrary nodes and arbitrary edges */
type Disjoint struct {
	Nodes NodeList
	Edges EdgeList
}

func NewDisjoint(nodes NodeList, edges EdgeList) *Disjoint {
	return &Disjoint{
		Nodes: nodes,
		Edges: edges,
	}
}

// ComposeSections takes a list of UI CDS graphs and composes them into a new single disjoint
func ComposeSections(graphs []*Section) *Disjoint {
	dj := &Disjoint{}

	// TODO: Create a disjoint from a list of sub-graphs, branches, other disjoints, etc.
	// 		add nodes to list and verify uniqueness of list (as each node is added)
	// 		add edges to map and verify each nodes (key) edge list is unique (as we are adding edges)
	// TODO: a possible optimization could be checking the uniqueness of each edge list ONCE after
	// 		all edges have been added, but ONLY for nodes that have had edges added to them,
	// 		and just create a new unique edge list for each of these nodes.
	return dj
}

func (d *Disjoint) ListNodes() NodeList {
	return d.Nodes
}

func (d *Disjoint) ListEdges() EdgeList {
	return d.Edges
}
