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

// FIXME: the problem with sections in general is that we are *not* using
// pointers to lists instead of actual lists. Which makes our complexity scale linearly.
type Section interface {
	ListNodes() NodeList
	ListEdges() EdgeList
	// NOTE: UpdateNodes and UpdateEdges can be used to add or remove nodes
	// and edges from the Sections list and provide the section with a new list.
	UpdateNodeList(NodeList)
	UpdateEdgeList(EdgeList)
}

/* Sub-graphs are non-disjoint collections of nodes and edges */
type Subgraph struct {
	Nodes NodeList
	Edges EdgeList
}

// NewSubgraph will grab all edges from nodes that connect to
// other nodes that are in our list.
func NewSubgraph(nodes NodeList, cp *CDS) *Subgraph {
	c := *cp
	edges := make(EdgeList, 0)

	for _, np := range nodes {
		n := *np
		cdsEdges := c.ListEdges()
		for _, ep := range cdsEdges {
			e := *ep
			sp := e.GetSource()
			dp := e.GetDestination()
			d := *dp
			if d.ID() == n.ID() && ContainsNode(nodes, sp) {
				edges = append(edges, ep)
			}
		}

	}

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

func (s *Subgraph) UpdateNodeList(n NodeList) {
	s.Nodes = n
}

func (s *Subgraph) UpdateEdgeList(e EdgeList) {
	s.Edges = e
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

func NewBranch(root *Node, cp *CDS) *Branch {
	edges := make(EdgeList, 0)
	nodes := make(NodeList, 0)

	nodes, edges = dfs(root, nodes, edges, cp)

	return &Branch{
		Nodes: nodes,
		Edges: edges,
	}
}

func dfs(start *Node, nodes NodeList, edges EdgeList, cp *CDS) (NodeList, EdgeList) {
	c := *cp
	// if node is not already in branch -- add
	if !ContainsNode(nodes, start) {
		nodes = append(nodes, start)
	}

	for _, ep := range c.ListEdges() {
		e := *ep
		// for all edges in CDS with node as source
		if e.GetSource() == start {
			if !ContainsEdge(edges, ep) {
				// add edge to branch
				edges = append(edges, ep)
				// for the destination node, add node and its edges to branch
				nodes, edges = dfs(e.GetDestination(), nodes, edges, cp)
			}
		}
	}

	return nodes, edges
}

func (b *Branch) ListNodes() NodeList {
	return b.Nodes
}

func (b *Branch) ListEdges() EdgeList {
	return b.Edges
}

func (b *Branch) UpdateNodeList(n NodeList) {
	b.Nodes = n
}

func (b *Branch) UpdateEdgeList(e EdgeList) {
	b.Edges = e
}

/*
	Partitions are only for linear CDSs
	(i.e. each node can only have at most 2 edges)
*/
type Partition struct {
	Nodes NodeList
	Edges EdgeList
}

func NewPartition(start, end *Node, cp *CDS) *Partition {
	nodes := make(NodeList, 0)
	edges := make(EdgeList, 0)

	nodes, edges = partDFS(start, end, nodes, edges, cp)

	return &Partition{
		Nodes: nodes,
		Edges: edges,
	}
}

func partDFS(startp, endp *Node, nodes NodeList, edges EdgeList, cp *CDS) (NodeList, EdgeList) {
	c := *cp
	start := *startp
	end := *endp
	// add node to partition nodes
	if !ContainsNode(nodes, startp) {
		nodes = append(nodes, startp)
		if start.ID() == end.ID() {
			return nodes, edges
		}
	}

	for _, ep := range c.ListEdges() {
		e := *ep
		// for all edges in CDS with node as source
		if e.GetSource() == startp {
			if !ContainsEdge(edges, ep) {
				// add edge to branch
				edges = append(edges, ep)
				// for the destination node, add node and its edge to branch
				nodes, edges = partDFS(e.GetDestination(), endp, nodes, edges, cp)
			}
		}
	}

	return nodes, edges
}

func (p *Partition) ListNodes() NodeList {
	return p.Nodes
}

func (p *Partition) ListEdges() EdgeList {
	return p.Edges
}

func (p *Partition) UpdateNodeList(n NodeList) {
	p.Nodes = n
}

func (p *Partition) UpdateEdgeList(e EdgeList) {
	p.Edges = e
}

/* Subsets are used for generic node selection (but not generic edge selection) */
type Subset struct {
	Nodes NodeList
	Edges EdgeList
}

// NewSubset grabs all (and only all) edges that are connected
// to a node in the list of nodes supplied.
func NewSubset(nodes NodeList, cp *CDS) *Subset {
	c := *cp
	cdsEdges := c.ListEdges()
	edges := make(EdgeList, 0)
	for _, n := range nodes {
		for _, ep := range cdsEdges {
			e := *ep
			if e.GetSource() == n || e.GetDestination() == n {
				if !ContainsEdge(edges, ep) {
					edges = append(edges, ep)
				}
			}
		}
	}

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

func (s *Subset) UpdateNodeList(n NodeList) {
	s.Nodes = n
}

func (s *Subset) UpdateEdgeList(e EdgeList) {
	s.Edges = e
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

// ComposeSections takes a list of CDS graphs (sections) and composes them into a new single disjoint
func ComposeSections(graphs []*Section) *Disjoint {
	nodes := make(NodeList, 0)
	edges := make(EdgeList, 0)

	for _, gp := range graphs {
		g := *gp
		gn := g.ListNodes()
		ge := g.ListEdges()

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
		Nodes: nodes,
		Edges: edges,
	}
}

func (d *Disjoint) ListNodes() NodeList {
	return d.Nodes
}

func (d *Disjoint) ListEdges() EdgeList {
	return d.Edges
}

func (d *Disjoint) UpdateNodeList(n NodeList) {
	d.Nodes = n
}

func (d *Disjoint) UpdateEdgeList(e EdgeList) {
	d.Edges = e
}
