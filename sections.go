package fabric

/* Sub-graphs are non-disjoint collections of nodes and edges */
type Subgraph struct {
	Nodes []int
	Edges map[int][]int
}

func NewSubgraph(nodes []int) *Subgraph {

	// TODO: will grab all edges from nodes that connect to other nodes that are in our list

	return &Subgraph{
		Nodes: nodes,
		Edges: make(map[int][]int),
	}
}

func (s *Subgraph) NodeCount() int {
	return len(s.Nodes)
}

func (s *Subgraph) EdgeCount() int {
	var total int
	for i, v := range s.Edges {
		total += len(v)
	}
	return total
}

/*
	Branches are all nodes and edges for a particuliar branch
	(usually of a tree graph)
	A branch is technically a sub-graph as well.
*/
type Branch struct {
	Nodes []int
	Edges map[int][]int
}

func NewBranch(root int) *Branch {
	// TODO: grab all children nodes recursively
	return &Branch{
		Nodes: nodes,
		Edges: edges,
	}
}

func (b *Branch) NodeCount() int {
	return len(b.Nodes)
}

func (b *Branch) EdgeCount() int {
	var total int
	for i, v := range b.Edges {
		total += len(v)
	}
	return total
}

/*
	Partitions are only for linear CDSs
	(i.e. each node can only have at most 2 edges)
*/
type Partition struct {
	Nodes []int
	Edges map[int][]int
}

func NewPartition(start, end int) *Partition {
	// TODO: adds all nodes between and including the start and end node
	// will also grab all edges for these nodes
	var nodes []int
	return &Partition{
		Nodes: nodes,
		Edges: make(map[int][]int),
	}
}

func (p *Partition) NodeCount() int {
	return len(p.Nodes)
}

func (p *Partition) EdgeCount() int {
	var total int
	for i, v := range p.Edges {
		total += len(v)
	}
	return total
}

/* Subsets are used for generic node selection (but not generic edge selection) */
type Subset struct {
	Nodes []int
	Edges map[int][]int
}

func NewSubset(nodes []int) *Subset {
	// TODO: grab all (and only all) edges that are connected to a node in the list of nodes supplied
	return &Subset{
		Nodes: nodes,
		Edges: make(map[int][]int),
	}
}

func (s *Subset) NodeCount() int {
	return len(s.Nodes)
}

func (s *Subset) EdgeCount() int {
	var total int
	for i, v := range s.Edges {
		total += len(v)
	}
	return total
}

/* Disjoints are a collection of arbitrary nodes and arbitrary edges */
type Disjoint struct {
	Nodes []int
	Edges map[int][]int
}

func NewDisjoint(nodes []int, edges map[int][]int) *Disjoint {
	return &Disjoint{
		Nodes: nodes,
		Edges: edges,
	}
}

// UICompose takes a list of UI CDS graphs and composes them into a new single disjoint
func UICompose(graphs []*UI) *Disjoint {
	// TODO: Create a disjoint from a list of sub-graphs, branches, other disjoints, etc.
	// 		add nodes to list and verify uniqueness of list (as each node is added)
	// 		add edges to map and verify each nodes (key) edge list is unique (as we are adding edges)
	// TODO: a possible optimization could be checking the uniqueness of each edge list ONCE after
	// 		all edges have been added, but ONLY for nodes that have had edges added to them,
	// 		and just create a new unique edge list for each of these nodes.
}

func (d *Disjoint) NodeCount() int {
	return len(d.Nodes)
}

func (d *Disjoint) EdgeCount() int {
	var total int
	for i, v := range d.Edges {
		total += len(v)
	}
	return total
}
