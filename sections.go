package fabric

/*
	Extensional Lists vs. Intensional Conditions

	Intensional Conditions: would have to be part of a Breadth-First or Depth-First traversal,
	where every node and edge in the CDS would be compared against the defined condition,
	and if it passed, would be considered accessible by the UI.

	Since these complete-traversal checks could be computationally expensive to perform every
	time we want to execute an access procedure, we instead will prefer the creation of extensional
	lists that are attached to a UI that can be used.

	Intensional Conditions can be used if the CDS is small enough to not become a processing burden.
*/

// Section is the interface definition for cutting out a section of the global CDS
type Section interface {
	ListNodes() *NodeList
	ListEdges() *EdgeList
	// NOTE: UpdateNodes and UpdateEdges can be used to add or remove nodes
	// and edges from the Sections list and provide the section with a new list.
	UpdateNodeList(*NodeList)
	UpdateEdgeList(*EdgeList)
}

/* Sub-graphs are non-disjoint collections of nodes and edges */

// Subgraph ...
type Subgraph struct {
	Nodes *NodeList
	Edges *EdgeList
}

// NewSubgraph will grab all edges from nodes that connect to
// other nodes that are in our list.
func NewSubgraph(nlp *NodeList, c CDS) Section {
	nodes := *nlp
	edges := make(EdgeList, 0)

	for _, n := range nodes {
		cdsEdges := c.ListEdges()
		for _, e := range cdsEdges {
			s := e.GetSource()
			d := e.GetDestination()
			if d.ID() == n.ID() && ContainsNode(nodes, s) {
				edges = append(edges, e)
			}
		}

	}

	return &Subgraph{
		Nodes: nlp,
		Edges: &edges,
	}
}

// ListNodes ...
func (s *Subgraph) ListNodes() *NodeList {
	return s.Nodes
}

// ListEdges ...
func (s *Subgraph) ListEdges() *EdgeList {
	return s.Edges
}

// UpdateNodeList ...
func (s *Subgraph) UpdateNodeList(nlp *NodeList) {
	s.Nodes = nlp
}

// UpdateEdgeList ...
func (s *Subgraph) UpdateEdgeList(elp *EdgeList) {
	s.Edges = elp
}

/*
	Branches are all nodes and edges for a particuliar branch
	(usually of a tree graph)
	A branch is technically a sub-graph as well.
*/

// Branch ...
type Branch struct {
	Nodes *NodeList
	Edges *EdgeList
}

// NewBranch ...
func NewBranch(root Node, c CDS) Section {
	edges := make(EdgeList, 0)
	nodes := make(NodeList, 0)

	nodes, edges = dfs(root, nodes, edges, c)

	return &Branch{
		Nodes: &nodes,
		Edges: &edges,
	}
}

func dfs(start Node, nodes NodeList, edges EdgeList, c CDS) (NodeList, EdgeList) {
	// if node is not already in branch -- add
	if !ContainsNode(nodes, start) {
		nodes = append(nodes, start)
	}

	for _, e := range c.ListEdges() {
		// for all edges in CDS with node as source
		if e.GetSource() == start {
			if !ContainsEdge(edges, e) {
				// add edge to branch
				edges = append(edges, e)
				// for the destination node, add node and its edges to branch
				nodes, edges = dfs(e.GetDestination(), nodes, edges, c)
			}
		}
	}

	return nodes, edges
}

// ListNodes ...
func (b *Branch) ListNodes() *NodeList {
	return b.Nodes
}

// ListEdges ...
func (b *Branch) ListEdges() *EdgeList {
	return b.Edges
}

// UpdateNodeList ...
func (b *Branch) UpdateNodeList(nlp *NodeList) {
	b.Nodes = nlp
}

// UpdateEdgeList ...
func (b *Branch) UpdateEdgeList(elp *EdgeList) {
	b.Edges = elp
}

/*
	Partitions are only for linear CDSs
	(i.e. each node can only have at most 2 edges)
*/

// Partition ...
type Partition struct {
	Nodes *NodeList
	Edges *EdgeList
}

// NewPartition ...
func NewPartition(start, end Node, c CDS) Section {
	nodes := make(NodeList, 0)
	edges := make(EdgeList, 0)

	nodes, edges = partDFS(start, end, nodes, edges, c)

	return &Partition{
		Nodes: &nodes,
		Edges: &edges,
	}
}

func partDFS(start, end Node, nodes NodeList, edges EdgeList, c CDS) (NodeList, EdgeList) {
	// add node to partition nodes
	if !ContainsNode(nodes, start) {
		nodes = append(nodes, start)
		if start.ID() == end.ID() {
			return nodes, edges
		}
	}

	for _, e := range c.ListEdges() {
		// for all edges in CDS with node as source
		if e.GetSource() == start {
			if !ContainsEdge(edges, e) {
				// add edge to branch
				edges = append(edges, e)
				// for the destination node, add node and its edge to branch
				nodes, edges = partDFS(e.GetDestination(), end, nodes, edges, c)
			}
		}
	}

	return nodes, edges
}

// ListNodes ...
func (p *Partition) ListNodes() *NodeList {
	return p.Nodes
}

// ListEdges ...
func (p *Partition) ListEdges() *EdgeList {
	return p.Edges
}

// UpdateNodeList ...
func (p *Partition) UpdateNodeList(nlp *NodeList) {
	p.Nodes = nlp
}

// UpdateEdgeList ...
func (p *Partition) UpdateEdgeList(elp *EdgeList) {
	p.Edges = elp
}

/* Subsets are used for generic node selection (but not generic edge selection) */

// Subset ...
type Subset struct {
	Nodes *NodeList
	Edges *EdgeList
}

// NewSubset grabs all (and only all) edges that are connected
// to a node in the list of nodes supplied.
func NewSubset(nlp *NodeList, c CDS) Section {
	nodes := *nlp
	cdsEdges := c.ListEdges()
	edges := make(EdgeList, 0)
	for _, n := range nodes {
		for _, e := range cdsEdges {
			if e.GetSource() == n || e.GetDestination() == n {
				if !ContainsEdge(edges, e) {
					edges = append(edges, e)
				}
			}
		}
	}

	return &Subset{
		Nodes: nlp,
		Edges: &edges,
	}
}

// ListNodes ...
func (s *Subset) ListNodes() *NodeList {
	return s.Nodes
}

// ListEdges ...
func (s *Subset) ListEdges() *EdgeList {
	return s.Edges
}

// UpdateNodeList ...
func (s *Subset) UpdateNodeList(nlp *NodeList) {
	s.Nodes = nlp
}

// UpdateEdgeList ...
func (s *Subset) UpdateEdgeList(elp *EdgeList) {
	s.Edges = elp
}

/* Disjoints are a collection of arbitrary nodes and arbitrary edges */

// Disjoint ...
type Disjoint struct {
	Nodes *NodeList
	Edges *EdgeList
}

// NewDisjoint ...
func NewDisjoint(nlp *NodeList, elp *EdgeList) *Disjoint {
	return &Disjoint{
		Nodes: nlp,
		Edges: elp,
	}
}

// ComposeSections takes a list of CDS graphs (sections) and composes them into a new single disjoint
func ComposeSections(graphs []*Section) Section {
	nodes := make(NodeList, 0)
	edges := make(EdgeList, 0)

	for _, gp := range graphs {
		g := *gp
		gnp := g.ListNodes()
		gn := *gnp
		gep := g.ListEdges()
		ge := *gep

		// add graph nodes to Disjoint node list
		for _, n := range gn {
			if !ContainsNode(nodes, n) {
				nodes = append(nodes, n)
			}
		}

		// add graph edges to disjoint edge list
		for _, e := range ge {
			if !ContainsEdge(edges, e) {
				edges = append(edges, e)
			}
		}

	}

	return &Disjoint{
		Nodes: &nodes,
		Edges: &edges,
	}
}

// ListNodes ...
func (d *Disjoint) ListNodes() *NodeList {
	return d.Nodes
}

// ListEdges ...
func (d *Disjoint) ListEdges() *EdgeList {
	return d.Edges
}

// UpdateNodeList ...
func (d *Disjoint) UpdateNodeList(nlp *NodeList) {
	d.Nodes = nlp
}

// UpdateEdgeList ...
func (d *Disjoint) UpdateEdgeList(elp *EdgeList) {
	d.Edges = elp
}
